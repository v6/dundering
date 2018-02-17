# dundering
  // , every day I'm pufflin'

Install Go, then Docker


    git clone https://github.com/v6/dundering.git
    cd dundering
    go build
    sudo docker build -t aloofhello .
    sudo docker run -p 1337:31337 aloofhello

In the above commands, `sudo` is only necessary for using Docker in Linux.
