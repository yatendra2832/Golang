const express = require('express')
const app = express()
const port = 3000

app.use(express.json())
app.use(express.urlencoded({ extended: true }))

app.get('/', (req, res) => {
    res.status(200).send("Welcome To CodeWithYatendra Learning Golang")
})

app.get('/get', (req, res) => {
    res.status(200).json({ message: "Hello From CodeWithYatendra in the Journey to golang" })
})

app.post('/post', (req, res) => {
    let myJson = req.body
    res.status(200).send(myJson)
})

app.post('/postForm', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})

app.listen(port, () => {
    console.log(`App is listening at port ${port}`)
})