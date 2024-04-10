  const AWS = require('aws-sdk');
  const ddb = new AWS.DynamoDB.DocumentClient();
      
  exports.handler = async function (event, context) {
    console.log("QUERY : ", event['queryStringParameters'])
    //! Change
    try {
      await ddb
        .put({
          TableName: process.env.USER_TABLE_ARN,
          Item: {
            connectionId: event.requestContext.connectionId,
          },
        })
        .promise();
    } catch (err) {
      return {
        statusCode: 500,
      };
    }
    return {
      statusCode: 200,
    };
    //! 
  };