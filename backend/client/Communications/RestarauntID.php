<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: communications.proto

namespace Communications;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>communications.RestarauntID</code>
 */
class RestarauntID extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string ID = 1;</code>
     */
    private $ID = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $ID
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Communications::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string ID = 1;</code>
     * @return string
     */
    public function getID()
    {
        return $this->ID;
    }

    /**
     * Generated from protobuf field <code>string ID = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setID($var)
    {
        GPBUtil::checkString($var, True);
        $this->ID = $var;

        return $this;
    }

}

