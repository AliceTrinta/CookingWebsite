#env GOOS=linux GOARCH=amd64 go build
#docker build -t mercurius:web-example .
#docker run -p 8080:8080 -d mercurius:web-example

FROM scratch

ADD web-example /
ADD conf/ /conf
ADD public/ /public
ADD locale/ /locale

CMD [ "/web-example" ]