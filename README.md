# dundering
  // , every day I'm pufflin'

Install Go, then Docker


    git clone https://github.com/v6/dundering.git
    cd dundering
    go build
    sudo docker build -t aloofhello .
    sudo docker run --name aloofhello -p 1337:31337 aloofhello
    sudo docker exec --interactive aloofhello ls
    sudo docker exec --interactive aloofhello go build
    sudo docker exec --interactive aloofhello ./dundering

In the above commands, `sudo` is only necessary for using Docker in Linux.

If the above has gone well, test out your shiny new insult delivery server by copying and pasting the following two URLs into your browser address bar  :

http://localhost:1337

http://localhost:1337/assign

## Troubleshooting

Make sure you are in the proper directory. If `go build` doesn't work, you may need to do some additional steps after installing Go.
