FROM node

WORKDIR /app

RUN apt-get update && apt-get install -y \
    libssl-dev

COPY package.json .

RUN npm install --force

COPY . .

EXPOSE 3000

CMD ["npm", "start"]