FROM node:12-alpine

# ENV NODE_ENV production

USER node
WORKDIR /home/node

ADD . /home/node/
RUN npm install && npm run build
ENV NODE_ENV production

RUN npm ci

CMD ["npm", "run", "start:prod"]
