<?php
class Router
{
  private $request;
  private $supportedHttpMethods = array(
    "GET",
    "POST"
  );
  function __construct(IRequest $request)
  {
   $this->request = $request;
   echo "Hi";
  }
  function __call($name, $args)
  {
    list($route, $method) = $args;
    if(!in_array(strtoupper($name), $this->supportedHttpMethods))
    {
      $this->invalidMethodHandler();
    }
    $this->{strtolower($name)}[$this->formatRoute($route)] = $method;
  }
  /**
   * Removes trailing forward slashes from the right of the route.
   * @param route (string)
   */
  private function formatRoute($route)
  {
    $result = rtrim($route, '/');
    if ($result === '')
    {
      return '/';
    }
    return $result;
  }
  private function invalidMethodHandler()
  {
    header("{$this->request->serverProtocol} 405 Method Not Allowed");
  }
  private function defaultRequestHandler()
  {
    header("{$this->request->serverProtocol} 404 Not Founds");
  }
  /**
   * Resolves a route
   */
  function resolve()
  {
    $methodDictionary = $this->{strtolower($this->request->requestMethod)};
    $formatedRoute = $this->formatRoute($this->request->requestUri);
    echo $this->request->requestUri;
    echo $this->request->requestMethod;
    var_dump($formatedRoute);
    $method = $methodDictionary[$formatedRoute];
    var_dump($method);
    if(is_null($method))
    {
      $this->defaultRequestHandler();
      return;
    }
    echo call_user_func_array($method, array($this->request));
  }
  function __destruct()
  {
    $this->resolve();
  }
}