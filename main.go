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
type Question struct {
	Id          int    `json:"id"`
	Title       string `json:"Title"`
	Detail      string `json:"Detail"`
	OwnQuestion string `json:"OwnQuestion"`
}
type Answer struct {
	AnsText  string `json:"AnsText"`
	Fullname string `json:"Fullname"`
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
func PathDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "detailsChat.html", gin.H{
		"content":   "This is an index page...",
		"resultLog": members,
	})
}

func PathshowChat(c *gin.Context) {
	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"kancha")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content":   err.Error(),
			"resultLog": members,
		})
		// returns nil on error

	}

	defer db.Close()
	results, err := db.Query("SELECT questions.QuestID, questions.QuestTitle, questions.QuestDetail, member.Fullname  FROM questions INNER JOIN member ON questions.QuestMemberID = member.Memid")
	var quests = []Question{}

	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content":   err.Error(),
			"resultLog": members,
		})

	} else {
		for results.Next() {
			quest := Question{}
			// create an instance of `Bird` and write the result of the current row into it
			if err := results.Scan(&quest.Id, &quest.Title, &quest.Detail, &quest.OwnQuestion); err != nil {
				log.Fatalf("could not scan row: %v", err)
			}
			// append the current instance to the slice of birds
			quests = append(quests, quest)
		}

		c.HTML(http.StatusOK, "showChat.html", gin.H{
			"question":  quests,
			"resultLog": members,
		})
	}

}
func PathInChat(c *gin.Context) {
	c.HTML(http.StatusOK, "increaseChat.html", gin.H{
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
			"content":   err.Error(),
			"resultLog": members,
		}) // returns nil on error
	}

	defer db.Close()
	results, err := db.Query("INSERT INTO questions (QuestTitle, QuestDetail, QuestMemberID)VALUES (?, ?, ?);", title, detail, Memberid)
	_ = results
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content":   err.Error(),
			"resultLog": members,
		})
	} else {
		results, err := db.Query("SELECT questions.QuestID, questions.QuestTitle, questions.QuestDetail, member.Fullname  FROM questions INNER JOIN member ON questions.QuestMemberID = member.Memid")
		var quests = []Question{}

		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"content":   err.Error(),
				"resultLog": members,
			})

		} else {
			for results.Next() {
				quest := Question{}
				// create an instance of `Bird` and write the result of the current row into it
				if err := results.Scan(&quest.Id, &quest.Title, &quest.Detail, &quest.OwnQuestion); err != nil {
					log.Fatalf("could not scan row: %v", err)
				}
				// append the current instance to the slice of birds
				quests = append(quests, quest)
			}

			c.HTML(http.StatusOK, "showChat.html", gin.H{
				"question":  quests,
				"resultLog": members,
			})
		}
	}

}

func Logout(c *gin.Context) {
	members = nil
	Memberid = 0
	c.HTML(http.StatusOK, "login.html", gin.H{
		"resultLog": members,
	})
}
func DetailChat(c *gin.Context) {
	var idChat = c.Param("Qid")
	db, err := sql.Open("mysql", "root"+":"+""+"@tcp(127.0.0.1:3306)/"+"kancha")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content":   err.Error(),
			"resultLog": members,
		})
		// returns nil on error

	}

	defer db.Close()
	Qquery, err := db.Query("SELECT questions.QuestID, questions.QuestTitle, questions.QuestDetail, member.Fullname  FROM questions INNER JOIN member ON questions.QuestMemberID = member.Memid WHERE QuestID = ?", idChat)
	Aquery, err := db.Query("SELECT Anstext, Fullname FROM `answers` INNER JOIN member ON answers.AnsmemberID = member.Memid WHERE QuestID = ? ORDER BY Ansid", idChat)
	var quests = []Question{}
	var answers = []Answer{}
	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content":   err.Error(),
			"resultLog": members,
		})

	} else {
		for Qquery.Next() {
			quest := Question{}
			// create an instance of `Bird` and write the result of the current row into it
			if err := Qquery.Scan(&quest.Id, &quest.Title, &quest.Detail, &quest.OwnQuestion); err != nil {
				log.Fatalf("could not scan row: %v", err)
			}
			// append the current instance to the slice of birds
			quests = append(quests, quest)
		}
		for Aquery.Next() {
			ans := Answer{}
			// create an instance of `Bird` and write the result of the current row into it
			if err := Aquery.Scan(&ans.AnsText, &ans.Fullname); err != nil {
				log.Fatalf("could not scan row: %v", err)
			}
			// append the current instance to the slice of birds
			answers = append(answers, ans)
		}

		c.HTML(http.StatusOK, "detailsChat.html", gin.H{
			"question":  quests,
			"answer":    answers,
			"resultLog": members,
		})
	}

}
func AddAnswer(c *gin.Context) {
	QuestId := c.PostForm("Qid")
	text := c.PostForm("text-response")

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
	results, err := db.Query("INSERT INTO answers (Anstext, AnsmemberID, QuestID)VALUES (?, ?, ?);", text, Memberid, QuestId)

	_ = results
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": err.Error(),
		})
	} else {
		Qquery, err := db.Query("SELECT questions.QuestID, questions.QuestTitle, questions.QuestDetail, member.Fullname  FROM questions INNER JOIN member ON questions.QuestMemberID = member.Memid WHERE QuestID = ?", QuestId)
		Aquery, err := db.Query("SELECT Anstext, Fullname FROM `answers` INNER JOIN member ON answers.AnsmemberID = member.Memid WHERE QuestID = ? ORDER BY Ansid", QuestId)
		var quests = []Question{}
		var answers = []Answer{}
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"content":   err.Error(),
				"resultLog": members,
			})

		} else {
			for Qquery.Next() {
				quest := Question{}
				// create an instance of `Bird` and write the result of the current row into it
				if err := Qquery.Scan(&quest.Id, &quest.Title, &quest.Detail, &quest.OwnQuestion); err != nil {
					log.Fatalf("could not scan row: %v", err)
				}
				// append the current instance to the slice of birds
				quests = append(quests, quest)
			}
			for Aquery.Next() {
				ans := Answer{}
				// create an instance of `Bird` and write the result of the current row into it
				if err := Aquery.Scan(&ans.AnsText, &ans.Fullname); err != nil {
					log.Fatalf("could not scan row: %v", err)
				}
				// append the current instance to the slice of birds
				answers = append(answers, ans)
			}

			c.HTML(http.StatusOK, "detailsChat.html", gin.H{
				"question":  quests,
				"answer":    answers,
				"resultLog": members,
			})
		}
	}

}
func searchQuestion(c *gin.Context) {
	Qname := c.PostForm("search")

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
	results, err := db.Query("SELECT questions.QuestID, questions.QuestTitle, questions.QuestDetail, member.Fullname  FROM questions INNER JOIN member ON questions.QuestMemberID = member.Memid WHERE QuestTitle LIKE '%" + Qname + "%'")

	var quests = []Question{}

	if err != nil {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": err.Error(),
		})

	} else {
		for results.Next() {
			quest := Question{}
			// create an instance of `Bird` and write the result of the current row into it
			if err := results.Scan(&quest.Id, &quest.Title, &quest.Detail, &quest.OwnQuestion); err != nil {
				log.Fatalf("could not scan row: %v", err)
			}
			// append the current instance to the slice of birds
			quests = append(quests, quest)
		}

		c.HTML(http.StatusOK, "showChat.html", gin.H{
			"question":  quests,
			"resultLog": members,
		})
	}

}

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("gosession", store))
	//router.UseRawPath = true
	//router.UnescapePathValues = false // set session

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/css", "./css") // this is loading css file
	router.Static("/img", "./img")
	router.Static("/templates", "./templates")

	router.GET("/", getIndex)
	router.GET("/index.html", getIndex)
	router.GET("/detailsChat.html", PathDetail)
	router.GET("/increaseChat.html", PathInChat)
	router.GET("/showChat.html", PathshowChat)
	router.GET("/register.html", PathRegister)
	router.GET("/login.html", PathLogin)
	router.POST("/register", PostRegister)
	router.POST("/login", PostLogin)
	router.GET("/logout", Logout)
	router.GET("/detailsChat.html/:Qid", DetailChat)
	router.POST("/searchQuestion", searchQuestion)
	router.POST("/addAnswer", AddAnswer)
	router.POST("/addQuestion", AddQuestion)

	router.Run(":8080")

}
