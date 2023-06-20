FROM node:18-alpine

COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.7.0 /lambda-adapter /opt/extensions/lambda-adapter

ENV PORT=4000 NODE_ENV=production

WORKDIR /opt/expenses

COPY ./.next/standalone ./

CMD ["node","server.js"]