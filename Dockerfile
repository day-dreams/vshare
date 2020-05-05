FROM jrottenberg/ffmpeg:4.1-ubuntu

RUN cat /proc/version
RUN ffmpeg -version
RUN apt install -y tree

ADD build build
#ENTRYPOINT pwd && ls && tree
ENTRYPOINT addr=0.0.0.0:8080 config=./build/config.json ./build/VShareServer
