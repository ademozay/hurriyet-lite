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
    <?php
        $id = $_GET['id'];
        $resp = file_get_contents('http://api/article?id=' . $id);
        $article = json_decode($resp);
        echo "<h2>$article->Title</h2>";
        echo html_entity_decode($article->Text);
        echo "<i><h5>$article->Editor</h5><i/>";
    ?>
</body>
</html>