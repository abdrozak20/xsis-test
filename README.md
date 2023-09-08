# Getting Started
Follow these instructions to setting up and running the project on your local machine for development and testing
purposes.

## Run The Project
1. Clone this repository: `git clone https://github.com/abdrozak20/xsis-test.git`  
2. Change to master branch `git checkout master`
3. Fill the environment variables in the `.env` file  
4. Install go dependencies
   ```bash
   go get
   ```
5. Start the app 
    ```bash 
    go run main.go
    ```  
  
## DATABASE
RUN THIS..

CREATE TABLE `movies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `description` text NULL,
  `rating` float NULL DEFAULT 0,
  `image` varchar(100) NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;
;

## Request
- GET MOVIES (GET)
curl --location 'localhost:3001/Movies'

- GET DETAIL MOVIE (GET)
curl --location 'localhost:3001/Movies/1'

- UPLOAD IMAGE (POST)
curl --location 'localhost:3001/upload-image' \
--form 'type="movies"' \
--form 'file=@"/C:/Users/PSSCOB006336/Pictures/tempsnip.png"'

- CREATE MOVIE (POST)
curl --location 'localhost:3001/Movies' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Sangkuriang",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry'\''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
    "rating": 9,
    "image": ""
}'

- UPDATE MOVIE (PATCH)
curl --location --request PATCH 'localhost:3001/Movies/4' \
--header 'Content-Type: application/json' \
--data '{
            "id": "4",
            "title": "Sangkuriang 2",
            "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry'\''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            "rating": 7.7,
            "image": "assets/movie/1694156724696-tempsnip.png"
        }'

- DELETE MOVIE (DELETE)
curl --location --request DELETE 'localhost:3001/Movies/7'

