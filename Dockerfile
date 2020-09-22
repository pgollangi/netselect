FROM gcr.io/distroless/static
COPY netselect /netselect
ENTRYPOINT ["/netselect"]
