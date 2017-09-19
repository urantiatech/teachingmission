package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {

	r := gin.Default()
	// router.GET("/someGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {

		// Send response
		c.JSON(http.StatusOK, gin.H{
			"css": gin.H{
				"style.css": "/css/style.css",
			},
			"nav": gin.H{
				"submenu1": gin.H{
					"item1": gin.H{
						"icon": "xyz",
						"uri":  "/xyz",
					},
					"item2": gin.H{
						"icon": "xyz",
						"uri":  "/xyz",
					},
				},
				"submenu2": gin.H{
					"item1": gin.H{
						"icon": "xyz",
						"uri":  "/xyz",
					},
					"item2": gin.H{
						"icon": "xyz",
						"uri":  "/xyz",
					},
				},
			},
			"user": gin.H{
				"userid": "sangeetk",
				"name":   "Sangeet Kumar",
				"role":   "administrator",
				"image":  "sangeetk.png",
				"notifications": gin.H{
					"mail": gin.H{
						"1": "latest mail title...",
						"2": "mail title...",
					},
					"notifications": gin.H{
						"mail": gin.H{
							"1": "latest mail title...",
							"2": "mail title...",
						},
						"chat": gin.H{
							"1": gin.H{
								"msg":    "chat message 1",
								"sender": "Sender 1",
							},
							"2": gin.H{
								"msg":    "chat message 2",
								"sender": "Sender 2",
							},
						},
					},
				},
			},
			"contentarea": gin.H{
				"quote": gin.H{
					"text":    "Quote from the light...",
					"teacher": "Teacher Ophelius",
				},
				"body": gin.H{
					"title": "title of the homepage",
					"text":  "text...",
				},
			},
		})
	})

	r.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "God is Love")

		c.JSON(http.StatusOK, gin.H{
			"nav": gin.H{
				"submenu1": gin.H{
					"item1": gin.H{
						"icon": "xyz",
						"uri":  "/xyz",
					},
					"item2": gin.H{
						"icon": "xyz",
						"uri":  "/xyz",
					},
				},
			},
			"content": gin.H{
				"banner": gin.H{
					"image": "banner1.jpg",
					"text":  "This is banner text",
				},
				"search": gin.H{
					"keywords": query,
					"results": gin.H{
						"1": gin.H{
							"transcriptid": "transcriptid",
							"lang":         "en",
							"title":        "Title of the transcript",
							"date":         "yyyy-mm-dd",
							"teachers": gin.H{
								"1": gin.H{
									"id":   "christ-michael",
									"name": "Christ Michael",
								},
								"2": gin.H{
									"id":   "monjoronson",
									"name": "Monjoronson",
								},
								"3": gin.H{
									"id":   "machiventa-melchizedek",
									"name": "Machiventa Melchizedek",
								},
							},
							"category": gin.H{
								"1": gin.H{
									"catid": "category-1",
									"name":  "Category 1",
								},
								"2": gin.H{
									"catid": "category-2",
									"name":  "Category 2",
								},
								"3": gin.H{
									"catid": "category-3",
									"name":  "Category 3",
								},
							},
							"text": "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua ...",
						},
						"2": gin.H{
							"transcriptid": "transcriptid",
							"lang":         "en",
							"title":        "Title of the transcript",
							"date":         "yyyy-mm-dd",
							"teachers": gin.H{
								"1": gin.H{
									"id":   "christ-michael",
									"name": "Christ Michael",
								},
								"2": gin.H{
									"id":   "monjoronson",
									"name": "Monjoronson",
								},
							},
							"category": gin.H{
								"1": gin.H{
									"catid": "category-2",
									"name":  "Category 2",
								},
							},
							"text": "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua ...",
						},
						"3": gin.H{
							"transcriptid": "transcriptid",
							"lang":         "en",
							"title":        "Title of the transcript",
							"date":         "yyyy-mm-dd",
							"teachers": gin.H{
								"1": gin.H{
									"id":   "machiventa-melchizedek",
									"name": "Machiventa Melchizedek",
								},
							},
							"category": gin.H{
								"1": gin.H{
									"catid": "category-1",
									"name":  "Category 1",
								},
							},
							"text": "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua ...",
						},
						"4": gin.H{
							"transcriptid": "transcriptid",
							"lang":         "en",
							"title":        "Title of the transcript",
							"date":         "yyyy-mm-dd",
							"teachers": gin.H{
								"1": gin.H{
									"id":   "christ-michael",
									"name": "Christ Michael",
								},
							},
							"category": gin.H{
								"1": gin.H{
									"catid": "category-3",
									"name":  "Category 3",
								},
							},
							"text": "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua ...",
						},
						"5": gin.H{
							"transcriptid": "transcriptid",
							"lang":         "en",
							"title":        "Title of the transcript",
							"date":         "yyyy-mm-dd",
							"teachers": gin.H{
								"1": gin.H{
									"id":   "christ-michael",
									"name": "Christ Michael",
								},
								"2": gin.H{
									"id":   "monjoronson",
									"name": "Monjoronson",
								},
							},
							"category": gin.H{
								"1": gin.H{
									"catid": "category-1",
									"name":  "Category 1",
								},
								"2": gin.H{
									"catid": "category-3",
									"name":  "Category 3",
								},
							},
							"text": "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua ...",
						},
					},
				},
			},
		})
	})

	r.GET("/transcript/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"id":            id,
			"date":          "2011-07-07",
			"group":         "11:11 Progress Group",
			"import_date":   "2017-04-24",
			"import_domain": "1111angels.net",
			"import_url":    "http://1111angels.net/old_files/E_Archives/list1579.htm",
			"language":      "en",
			"paragraphs": gin.H{
				"1":  "Urantia, July 7, 2011.",
				"2":  "Teacher Monjoronson.",
				"3":  "Received by Lytske.",
				"4":  "Monjoronson: \"It is just about the right time for you and I to make our acquaintance, child. I have been waiting a long time 'in the wings', so to speak, to find an entrance into your mind, and a measure of trust and willingness on your part, to hear what I have to say.",
				"5":  "\"It is so very important that more, and increasingly more people become aware as to who I am, and what I will be about in the near future here on this lovely, but oh-so-backward planet, where some minds have not emerged into the Light and the realization that they, too, are children of God -- His created but evolving beings destined to a grand future, if they learn to extend mercy and love to each other.",
				"6":  "\"As long as fear runs rampant, and a segment of your population still lives in cave times -- meaning that they have not come to the realization that there is nothing to defend themselves against -- where enemies are a figment of their imagination.",
				"7":  "\"It is this tendency to barbarism and war why this benighted planet is stuck in her evolution, and in grave danger of sliding backward. Her lovely face is marred by the goings on with offensive weaponry, which is getting more devastating as time goes on.",
				"8":  "\"It is because of this war-like propensity, and you not having learned respect for human life, that the celestial powers are creating wake-up calls to all who can hear this call for mercy towards one another.",
				"9":  "\"The time is now to start behaving like responsible adults instead of little children in a sandbox throwing sand in each other's eyes. You have become more war-like in your evolution rather than growing more sensible and loving.",
				"10": "\"Where is the balance and harmony in your lives? Do you ever think about this? Surely you also must yearn for inner peace and rest in your souls?",
				"11": "\"You are on the threshold of a wonderful era of peace, but I do need, as another Son of God, your cooperation to make our plans manifest, and turn this beloved planet into a garden jewel of creation.",
				"12": "\"Therefore, I am asking for your help. Please, all those who will chance to read these words, think of how you can cooperate in becoming more loving and peaceful yourselves. Surely, there is no need to lord it over others. People all carry the same responsibility for growing into the divine blueprint the Creator God has bestowed upon each of you at birth.",
				"13": "\"Know you not that you all hail from the same root Source and are therefore kin? Extend mercy and love towards all and learn how to live a God-guided life.",
				"14": "\"Thank you.\"",
				"15": "© 11:11 Progress Group.",
				"16": "You lit a Flame and it will become a Raging Fire - ABC-22.",
				"17": "www.1111angels.com",
			},
			"spider": "monjoronson",
			"status": "live",
			"teacher": gin.H{
				"1": "Teacher Monjoronson",
			},
			"text": gin.H{
				"1": "<p class=\"small\" align=\"left\">Urantia, July 7, 2011. <br>\n          <b>Teacher Monjoronson. </b><br>\n          \n          Received by Lytske. <br><br>\n          <b>Monjoronson</b>: \"It is just about the right time for you and \n          I to make our acquaintance, child. I have been waiting a long time 'in \n          the wings', so to speak, to find an entrance into your mind, and a measure \n          of trust and willingness on your part, to hear what I have to say. <br>\n          <br>\n          \"It is so very important that more, and increasingly more people \n          become aware as to who I am, and what I will be about in the near future \n          here on this lovely, but oh-so-backward planet, where some minds have \n          not emerged into the Light and the realization that they, too, are children \n          of God -- His created but evolving beings destined to a grand future, \n          if they learn to extend mercy and love to each other. <br>\n          <br>\n          \"As long as fear runs rampant, and a segment of your population \n          still lives in cave times -- meaning that they have not come to the \n          realization that there is nothing to defend themselves against -- where \n          enemies are a figment of their imagination. <br><br>\n          \"It is this tendency to barbarism and war why this benighted planet \n          is stuck in her evolution, and in grave danger of sliding backward. \n          Her lovely face is marred by the goings on with offensive weaponry, \n          which is getting more devastating as time goes on. <br><br>\n          \"It is because of this war-like propensity, and you not having \n          learned respect for human life, that the celestial powers are creating \n          wake-up calls to all who can hear this call for mercy towards one another. \n          <br><br>\n          \"The time is now to start behaving like responsible adults instead \n          of little children in a sandbox throwing sand in each other's eyes. \n          You have become more war-like in your evolution rather than growing \n          more sensible and loving. <br><br>\n          \"Where is the balance and harmony in your lives? Do you ever think \n          about this? Surely you also must yearn for inner peace and rest in your \n          souls? <br><br>\n          \"You are on the threshold of a wonderful era of peace, but I do \n          need, as another Son of God, your cooperation to make our plans manifest, \n          and turn this beloved planet into a garden jewel of creation. <br><br>\n          \"Therefore, I am asking for your help. Please, all those who will \n          chance to read these words, think of how you can cooperate in becoming \n          more loving and peaceful yourselves. Surely, there is no need to lord \n          it over others. People all carry the same responsibility for growing \n          into the divine blueprint the Creator God has bestowed upon each of \n          you at birth. <br><br>\n          \"Know you not that you all hail from the same root Source and are \n          therefore kin? Extend mercy and love towards all and learn how to live \n          a God-guided life. <br><br>\n          \"Thank you.\" \n        </p>",
				"2": "<p class=\"small\" align=\"center\"><b>© 11:11 Progress Group.</b><br>\n          <br>\n          <b><font color=\"#FF0000\">You lit a Flame and it will become a Raging \n          Fire - ABC-22.<br>\n          </font></b></p>",
				"3": "<p class=\"small\" align=\"center\"> <a href=\"http://www.1111angels.com\">www.1111angels.com</a> \n        </p>",
			},
			"title":      "Mercy and Love.",
			"updated_on": "2017-04-24",
			"visibility": "show",
		})
	})

	r.GET("/quote/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"teacher": "Teacher Ophelius",
			"text":    "It is the Father’s Will that this planet is to be healed of its dark past and brought into the age of Light and Life, and so those who are willfully thinking and doing those things that move this world ever so much closer to that enlightened era are doing the Will of the Father.",
		})
	})

	r.GET("/banner", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"url":  "/images/banner1.jpg",
			"text": "This is banner text.",
		})
	})

	r.Run(":8081")
}
