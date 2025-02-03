CREATE MIGRATION m1q7rf4ymlxqafa7zq7pnxtxc6yfpspvjrtqqmxm3qaxo5kjab72lq
    ONTO initial
{
  CREATE TYPE default::User {
      CREATE REQUIRED PROPERTY age: std::int32;
      CREATE REQUIRED PROPERTY name: std::str;
  };
};
