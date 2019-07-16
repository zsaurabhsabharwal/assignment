<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Communications;

/**
 */
class addeditClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Communications\RestarauntInfo $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function AddRestaraunt(\Communications\RestarauntInfo $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/communications.addedit/AddRestaraunt',
        $argument,
        ['\Communications\RestarauntID', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Communications\RestarauntInfo $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function EditRestaraunt(\Communications\RestarauntInfo $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/communications.addedit/EditRestaraunt',
        $argument,
        ['\Communications\RestarauntID', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Communications\Response $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetRestaraunt(\Communications\Response $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/communications.addedit/GetRestaraunt',
        $argument,
        ['\Communications\ListRestaraunt', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Communications\RestarauntID $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function DeleteRestaraunt(\Communications\RestarauntID $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/communications.addedit/DeleteRestaraunt',
        $argument,
        ['\Communications\Response', 'decode'],
        $metadata, $options);
    }

}
