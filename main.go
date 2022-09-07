package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	Id       int    `json:"id"`
	Fullname string `json:"first_name"`
	Email    string `json:"last_name"`
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"resultLog": members,
	})
}
func PathLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "This is an index page...",
	})
}
func PathRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"content": "This is an index page...",
	})
}

func PathshowChat(c *gin.Context) {
	c.HTML(http.StatusOK, "showChat.html", gin.H{
		"resultLog": members,
	})
}
func PathInChat(c *gin.Context) {
	c.HTML(http.StatusOK, "increaseChat.html", gin.H{
		"resultLog": members,
	})
}
func PathDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "detailsChat.html", gin.H{
		"resultLog": members,
	})
}

// Login
func PostRegister(c *gin.Context) {
	fullname := c.PostForm("fullname")
	email := c.PostForm("email")
	password := c.PostForm("password")

	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"kancha")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		c.HTML(http.StatusOK, "register.html", gin.H{
			"content": err.Error(),
		})
		// returns nil on error

	}

	defer db.Close()
	results, err := db.Query("INSERT INTO member (Fullname, Email, Pwd)VALUES (?, ?, ?);", fullname, email, password)

	if err != nil {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"content": err.Error(),
		})
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"resultReg": results,
		})
	}

}

var members = []Member{}
var Memberid = 0

func PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"kancha")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		c.HTML(http.StatusOK, "register.html", gin.H{
			"content": err.Error(),
		})
		// returns nil on error

	}

	defer db.Close()
	results, err := db.Query("SELECT Memid, Fullname, Email FROM member WHERE Email = ? AND Pwd = ?", email, password)

	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": err.Error(),
		})
	} else {
		for results.Next() {
			member := Member{}
			// create an instance of `Bird` and write the result of the current row into it
			if err := results.Scan(&member.Id, &member.Fullname, &member.Email); err != nil {
				log.Fatalf("could not scan row: %v", err)
			}
			// append the current instance to the slice of birds
			members = append(members, member)
			Memberid = member.Id
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"resultLog": members,
		})
	}

}
func AddQuestion(c *gin.Context) {
	title := c.PostForm("title")
	detail := c.PostForm("details")

	if members == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "This is an index page...",
		})
	}

	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"kancha")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		c.HTML(http.StatusOK, "register.html", gin.H{
			"content": err.Error(),
		})
		// returns nil on error

	}

	defer db.Close()
	results, err := db.Query("INSERT INTO questions (QuestTitle, QuestDetail, QuestMemberID)VALUES (?, ?, ?);", title, detail, Memberid)

	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": err.Error(),
		})
	} else {
		c.HTML(http.StatusOK, "showChat.html", gin.H{
			"content": results,
		})
	}

}

func getProperty(member *[]Member, s string) {
	panic("unimplemented")
}
func Logout(c *gin.Context) {
	members = nil
	c.HTML(http.StatusOK, "login.html", gin.H{
		"resultLog": members,
	})
}

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("gosession", store)) // set session

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/css", "./css") // this is loading css file
	router.Static("/img", "./img")
	router.Static("/templates", "./templates")

	router.GET("/", getIndex)
	router.GET("/index.html", getIndex)
	router.GET("/increaseChat.html", PathInChat)
	router.GET("/showChat.html", PathshowChat)
	router.GET("/register.html", PathRegister)
	router.GET("/login.html", PathLogin)
	router.GET("/detailsChat.html", PathDetail)
	router.POST("/register", PostRegister)
	router.POST("/login", PostLogin)
	router.GET("/logout", Logout)
	router.POST("/addQuestion", AddQuestion)

	router.Run(":8080")

}
