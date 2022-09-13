FROM scratch

WORKDIR /src/gin_blog

COPY . /src/gin_blog

EXPOSE 8000

CMD ["./gin_blog"]
