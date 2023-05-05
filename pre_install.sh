install_py(){
  sudo apt update && sudo apt install python3 && sudo apt install python3-pip -y
  sudo pip3 install pytest
  sudo pytest --version
}
install_py


BINARY="go1.19.5.linux-amd64.tar.gz"
install_go(){
    curl -OL "https://go.dev/dl/$BINARY"
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.19.5.linux-amd64.tar.gz
    rm $BINARY
    mkdir -p $HOME/go/bin
    echo "export PATH=\$PATH:\$HOME/go/bin" >> $HOME/.bashrc
    echo "export PATH=\$PATH:/usr/local/go/bin" >> $HOME/.bashrc
    . $HOME/.bashrc
    /usr/local/go/bin/go version
    echo " "
    echo "source ~/.bashrc"
}
install_go