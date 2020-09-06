#env GOOS=linux GOARCH=amd64 go build
#docker build -t mercurius:cooking-website .
#docker run -p 8080:8080 -d mercurius:cooking-website

FROM scratch

ADD cooking-website /
ADD conf/ /conf
ADD public/ /public
ADD locale/ /locale

CMD [ "/cooking-website" ]