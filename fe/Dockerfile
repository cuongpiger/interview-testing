FROM node:16
 
WORKDIR /app
# install app dependencies
COPY package.json /app
 
 
# add app
COPY . /app
RUN npm install
EXPOSE 3000
 
# start app
CMD ["npm", "start"]