package main

import (
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/f4814/iscraper/models"
	log "github.com/sirupsen/logrus"
)

type DBHelper struct {
	Client driver.Client
	DB     driver.Database

	GraphOptions driver.CreateGraphOptions
	Graph        driver.Graph

	Users    driver.Collection
	Items    driver.Collection
	Comments driver.Collection

	EdgeLikes   driver.Collection
	EdgeTags    driver.Collection
	EdgeChild   driver.Collection
	EdgeFollows driver.Collection
	EdgePosts   driver.Collection
}

func NewDBHelper(client driver.Client, dbName string) DBHelper {
	h := DBHelper{Client: client}

	var (
		err    error
		exists bool
	)

	// Initialize Database
	var db driver.Database
	if exists, err = client.DatabaseExists(nil, dbName); err != nil {
		log.Fatal(err)
	}

	if !exists {
		log.WithFields(log.Fields{"database": dbName}).Debug("Creating new Database")
		db, err = client.CreateDatabase(nil, dbName, nil)
	} else {
		log.WithFields(log.Fields{"database": dbName}).Debug("Loading existing Database")
		db, err = client.Database(nil, dbName)
	}

	if err != nil {
		log.Fatal(err)
	}
	h.DB = db

	// Initialize Graph
	var graph driver.Graph
	if exists, err = h.DB.GraphExists(nil, "instagram"); err != nil {
		log.Fatal(err)
	}

	if !exists {
		log.WithFields(log.Fields{
			"database": dbName,
			"graph":    "instagram",
		}).Debug("Creating new Graph")
		graph, err = h.DB.CreateGraph(nil, "instagram", nil) // TODO Opts
	} else {
		log.WithFields(log.Fields{
			"database": dbName,
			"graph":    "instagram",
		}).Debug("Loading Existing graph")
		graph, err = h.DB.Graph(nil, "instagram")
	}

	if err != nil {
		log.Fatal(err)
	}
	h.Graph = graph

	// Get Edge and vertex collections
	h.Users = h.initVertex("users")
	h.Items = h.initVertex("items")
	h.Comments = h.initVertex("comments")

	h.EdgeLikes = h.initEdge("edge_likes")
	h.EdgeTags = h.initEdge("edge_tags")
	h.EdgeChild = h.initEdge("edge_child")
	h.EdgeFollows = h.initEdge("edge_follows")
	h.EdgePosts = h.initEdge("edge_posts")

	// Initialize Indexes
	uniqueOpts := driver.EnsureHashIndexOptions{Unique: true}
	h.initHashIndex(h.Users, []string{"id", "username"}, &uniqueOpts)
	h.initHashIndex(h.Items, []string{"id"}, &uniqueOpts)
	h.initHashIndex(h.Comments, []string{"id", "pk"}, &uniqueOpts)
	h.initHashIndex(h.EdgeLikes, []string{"_from", "_to"}, &uniqueOpts)
	h.initHashIndex(h.EdgePosts, []string{"_from", "_to"}, &uniqueOpts)
	h.initHashIndex(h.EdgeFollows, []string{"_from", "_to"}, &uniqueOpts)
	h.initHashIndex(h.EdgeTags, []string{"_from", "_to"}, &uniqueOpts)
	h.initHashIndex(h.EdgeChild, []string{"_from", "_to"}, &uniqueOpts)

	return h
}

func (h *DBHelper) initVertex(name string) driver.Collection {
	var (
		c      driver.Collection
		err    error
		exists bool
	)

	if exists, err = h.Graph.VertexCollectionExists(nil, name); err != nil {
		log.Fatal(err)
	}

	if !exists {
		log.WithFields(log.Fields{
			"database":   h.DB.Name(),
			"graph":      h.Graph.Name(),
			"collection": name,
		}).Debug("Creating new Vertex Collection")
		c, err = h.Graph.CreateVertexCollection(nil, name)
	} else {
		log.WithFields(log.Fields{
			"database":   h.DB.Name(),
			"graph":      h.Graph.Name(),
			"collection": name,
		}).Debug("Loading existing Vertex Collection")
		c, err = h.Graph.VertexCollection(nil, name)
	}

	if err != nil {
		log.Fatal(err)
	}

	return c

}

func (h *DBHelper) initEdge(name string) driver.Collection {
	defs := []driver.EdgeDefinition{
		driver.EdgeDefinition{
			Collection: "edge_likes",
			From:       []string{"users"},
			To:         []string{"users"},
		},
		driver.EdgeDefinition{
			Collection: "edge_tags",
			From:       []string{"items"},
			To:         []string{"users"},
		},
		driver.EdgeDefinition{
			Collection: "edge_child",
			From:       []string{"items", "comments"},
			To:         []string{"comments"},
		},
		driver.EdgeDefinition{
			Collection: "edge_follows",
			From:       []string{"users"},
			To:         []string{"users"},
		},
		driver.EdgeDefinition{
			Collection: "edge_posts",
			From:       []string{"users"},
			To:         []string{"items"},
		},
	}

	var constraints driver.VertexConstraints
	for _, d := range defs {
		if d.Collection == name {
			constraints.From = d.From
			constraints.To = d.To
			break
		}
	}

	var (
		err    error
		c      driver.Collection
		exists bool
	)

	if exists, err = h.Graph.EdgeCollectionExists(nil, name); err != nil {
		log.Fatal(err)
	}

	if !exists {
		log.WithFields(log.Fields{
			"database":   h.DB.Name(),
			"graph":      h.Graph.Name(),
			"collection": name,
		}).Debug("Creating new Edge Collection")
		c, err = h.Graph.CreateEdgeCollection(nil, name, constraints)
	} else {
		log.WithFields(log.Fields{
			"database":   h.DB.Name(),
			"graph":      h.Graph.Name(),
			"collection": name,
		}).Debug("Loading existing Edge Collection")
		c, _, err = h.Graph.EdgeCollection(nil, name)
	}

	if err != nil {
		log.Fatal(err)
	}

	return c
}

func (h *DBHelper) initHashIndex(coll driver.Collection, fields []string,
	opts *driver.EnsureHashIndexOptions) {
	_, n, err := coll.EnsureHashIndex(nil, fields, opts)
	if err != nil {
		log.Fatal(err)
	} else if n {
		log.WithFields(log.Fields{
			"database":   h.DB.Name(),
			"collection": coll.Name(),
			"fields":     fields,
			"opts":       opts,
		}).Debug("Created new Hash Index")
	} else {
		log.WithFields(log.Fields{
			"database":   h.DB.Name(),
			"collection": coll.Name(),
			"fields":     fields,
			"opts":       opts,
		}).Debug("Loaded existing Hash Index")
	}
}

func (h *DBHelper) SaveUser(user *models.User) {
	query := "FOR u IN users FILTER u.id == @id RETURN u"
	cur, err := h.DB.Query(nil, query,
		map[string]interface{}{
			"id": user.ID,
		},
	)
	defer cur.Close()
	if err != nil {
		log.Fatal(err)
	}

	var (
		meta driver.DocumentMeta
		msg  string
	)
	if cur.HasMore() {
		var oldUser models.User
		if meta, err = cur.ReadDocument(nil, &oldUser); err != nil {
			log.Fatal(err)
		} else {
			if oldUser.ScrapedAt.After(user.ScrapedAt) {
				user.ScrapedAt = oldUser.ScrapedAt
			}

			meta, err = h.Users.UpdateDocument(nil, meta.Key, *user)
			msg = "Loaded Item Metadata"
		}
	} else {
		meta, err = h.Users.CreateDocument(nil, *user)
		msg = "Saved user"
	}

	if err != nil {
		log.Fatal(err)
	}
	user.SetMeta(meta)

	log.WithFields(log.Fields{
		"username": user.Username,
	}).Trace(msg)
}

func (h *DBHelper) SaveItem(item *models.Item) {
	query := fmt.Sprintf("FOR i IN items FILTER i.id == '%s' RETURN i", item.ID)
	cur, err := h.DB.Query(nil, query, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close()

	var (
		meta driver.DocumentMeta
		msg  string
	)
	if cur.HasMore() {
		var oldItem models.Item
		if meta, err = cur.ReadDocument(nil, &oldItem); err != nil {
			log.Fatal(err)
		} else {
			if oldItem.ScrapedAt.After(item.ScrapedAt) {
				item.ScrapedAt = oldItem.ScrapedAt
			}

			msg = "Loaded Item Metadata"
			meta, err = h.Items.UpdateDocument(nil, meta.Key, *item)
		}
	} else {
		meta, err = h.Items.CreateDocument(nil, *item)
		msg = "Saved Item"
	}

	if err != nil {
		log.Fatal(err)
	}
	item.SetMeta(meta)

	log.WithFields(log.Fields{
		"id": item.ID,
	}).Trace(msg)
}

func (h *DBHelper) UserFollows(user *models.User, follows *models.User) {
	query := "FOR e IN edge_follows FILTER e._from == @from && e._to == @to RETURN e"
	cur, err := h.DB.Query(nil, query,
		map[string]interface{}{
			"from": user.GetMeta().ID,
			"to":   follows.GetMeta().ID,
		},
	)
	defer cur.Close()

	if err != nil {
		log.Fatal(err)
	}

	if !cur.HasMore() {
		edge := models.Follows{
			From: string(user.GetMeta().ID),
			To:   string(follows.GetMeta().ID),
		}

		if _, err := h.EdgeFollows.CreateDocument(nil, edge); err != nil {
			log.Fatal(err)
		}

		log.WithFields(log.Fields{
			"username": user.Username,
			"follows":  follows.Username,
			"edge":     edge,
		}).Trace("Add Follow Edge")
	}
}

func (h *DBHelper) UserFollowed(user *models.User, followed *models.User) {
	h.UserFollows(followed, user)
}

func (h *DBHelper) UserPosts(user *models.User, item *models.Item) {
	query := "FOR e IN edge_posts FILTER e._from == @from && e._to == @to RETURN e"
	cur, err := h.DB.Query(nil, query,
		map[string]interface{}{
			"from": user.GetMeta().ID,
			"to":   item.GetMeta().ID,
		},
	)
	defer cur.Close()

	if err != nil {
		log.Fatal(err)
	}

	if !cur.HasMore() {
		edge := models.Posts{
			From: string(user.GetMeta().ID),
			To:   string(item.GetMeta().ID),
		}

		if _, err := h.EdgePosts.CreateDocument(nil, edge); err != nil {
			log.Fatal(err)
		}

		log.WithFields(log.Fields{
			"username": user.Username,
			"item":     item.ID,
			"edge":     edge,
		}).Trace("Add Posts edge")
	}
}

func (h *DBHelper) UserLikes(user *models.User, item *models.Item) {
	query := "FOR e IN edge_likes FILTER e._from == @from && e._to == @to RETURN e"
	cur, err := h.DB.Query(nil, query,
		map[string]interface{}{
			"from": user.GetMeta().ID,
			"to":   item.GetMeta().ID,
		},
	)
	defer cur.Close()

	if err != nil {
		log.Fatal(err)
	}

	if !cur.HasMore() {
		edge := models.Likes{
			From:       string(user.GetMeta().ID),
			To:         string(item.GetMeta().ID),
			IsTopliker: false,
		}

		if _, err := h.EdgeLikes.CreateDocument(nil, edge); err != nil {
			log.Fatal(err)
		}

		log.WithFields(log.Fields{
			"username": user.Username,
			"item":     item.ID,
			"edge":     edge,
		}).Trace("Add Likes edge")
	}
}

// Convenience Frontent for DBHelper.ItemTags so it can be used as relation
func (h *DBHelper) UserTagged(user *models.User, item *models.Item) {
	h.ItemTags(item, user, models.Tags{})
}

func (h *DBHelper) ItemTags(item *models.Item, user *models.User,
	edge models.Tags) {
	query := "FOR e IN edge_tags FILTER e._from == @from && e._to == @to RETURN e"
	cur, err := h.DB.Query(nil, query,
		map[string]interface{}{
			"from": item.GetMeta().ID,
			"to":   user.GetMeta().ID,
		},
	)
	defer cur.Close()

	if err != nil {
		log.Fatal(err)
	}

	if !cur.HasMore() {
		edge.From = string(item.GetMeta().ID)
		edge.To = string(user.GetMeta().ID)

		if _, err := h.EdgeTags.CreateDocument(nil, edge); err != nil {
			log.Fatal(err)
		}

		log.WithFields(log.Fields{
			"username": user.Username,
			"item":     item.ID,
			"edge":     edge,
		}).Trace("Add Tags edge")
	}
}

