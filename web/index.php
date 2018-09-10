<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="shortcut icon" href="http://www.hurriyet.com.tr/favicon.ico">
    <link rel="stylesheet" href="style.css">
    <title>Hürriyet Lite</title>
</head>
<body>
    <strong>
        <a class="heading" href="/">Hürriyet Lite</a>
    </strong>
    <hr/>
    <ul>
        <?php
            $resp = file_get_contents('http://api/articles');
            $articles = json_decode($resp);
            foreach ($articles as $article) {
                echo "
                    <li>
                        <a href='article?id=$article->Id'>$article->Title</a>
                    </li>";
            }
        ?>
    </ul>
</body>
</html>