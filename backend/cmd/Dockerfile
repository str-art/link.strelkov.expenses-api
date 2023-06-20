FROM golang:1.20 as build

WORKDIR /opt/expenses

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o request-handler ./main.go

FROM public.ecr.aws/lambda/go:1

COPY --from=build /opt/expenses/request-handler ${LAMBDA_TASK_ROOT}

CMD ["request-handler"]