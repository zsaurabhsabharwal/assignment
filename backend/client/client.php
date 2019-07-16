<?php

require dirname(__FILE__).'/vendor/autoload.php';


include_once dirname(__FILE__).'/Request.php';
include_once dirname(__FILE__).'/Router.php';
@include_once dirname(__FILE__).'/Communications/addeditClient.php';
@include_once dirname(__FILE__).'/Communications/RestarauntInfo.php';
@include_once dirname(__FILE__).'/Communications/RestarauntID.php';
@include_once dirname(__FILE__).'/Communications/Response.php';
@include_once dirname(__FILE__).'/Communications/ListRestaraunt.php';
@include_once dirname(__FILE__).'/GPBMetadata/Communications.php';


// $router = new Router(new Request);

$client = new Communications\addeditClient('localhost:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

function addRestaraunt($name, $rating, $cuisines, $address, $time, $cft, $ID, $client)
{

    $request = new Communications\RestarauntInfo();
    
    $request->setName($name);
    $request->setRating($rating);
    $request->setCuisines($cuisines);
    $request->setAddress($address);
    $request->setTime($time);
    $request->setCft($cft);
    $request->setID($ID);

    list($reply, $status) = $client->AddRestaraunt($request)->wait();
    $message = $reply->getID();

    return $message;
}

function editRestaraunt($name, $rating, $cuisines, $address, $time, $cft, $ID, $client)
{

    $request = new Communications\RestarauntInfo();
    
    $request->setName($name);
    $request->setRating($rating);
    $request->setCuisines($cuisines);
    $request->setAddress($address);
    $request->setTime($time);
    $request->setCft($cft);
    $request->setID($ID);

    list($reply, $status) = $client->EditRestaraunt($request)->wait();
    $message = $reply->getID();

    return $message;
}

function getRestaraunt($stat, $client)
{

    $request = new Communications\Response();
    $request->setStatus($stat);
    list($reply, $status) = $client->GetRestaraunt($request)->wait();

    // var_dump($reply->getList());
    $List = $reply->getList();
    
    for($i=0; $i<count($List); $i++){    
        echo $List[$i]->getName()." ".$List[$i]->getID()."\n";
    }    

    $message = "OK";

    return $message;
}

function deleteRestaraunt($ID, $client)
{

    $request = new Communications\RestarauntID();
    $request->setID($ID);
    list($reply, $status) = $client->DeleteRestaraunt($request)->wait();
    $message = $reply->getStatus();

    return $message;
}

$name = "Saurabh";
$rating = "3.9";
$cuisines = "Cafe, Fast food...";
$address = "Gurgaon";
$time = "6am";
$cft = "230 for two person";
$ID = '1';

// $router->get('/add', function($request) {
//     return "Added";
// });

echo addRestaraunt($name, $rating, $cuisines, $address, $time, $cft, $ID, $client)."\n";

$name = "Saurabh Sabharwal2";
$rating = "3.9";
$cuisines = "Cafe, Fast food...";
$address = "Gurgaon";
$time = "6am";
$cft = "230 for two person";
$ID = '2';

echo addRestaraunt($name, $rating, $cuisines, $address, $time, $cft, $ID, $client)."\n";

$name = "Saurabh Sabharwal";
$rating = "3.9";
$cuisines = "Cafe, Fast food...";
$address = "Gurgaon";
$time = "6am";
$cft = "230 for two person";
$ID = '1';

echo editRestaraunt($name, $rating, $cuisines, $address, $time, $cft, $ID, $client)."\n";

$stat = "OK";

echo getRestaraunt($ID, $client)."\n";

$ID = '1';

echo deleteRestaraunt($ID, $client)."\n";

$ID = '2';

echo deleteRestaraunt($ID, $client)."\n";



