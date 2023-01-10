FROM node:16

WORKDIR /app

COPY dist/ .

RUN npm install --force

ENV BASE_PATH=/permata-aksesoris

CMD ["npm", "start", "--", "-p", "3001"]