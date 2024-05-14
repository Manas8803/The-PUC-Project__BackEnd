const AWS = require("aws-sdk");
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async function (event, context) {
	const rtoOfficeName = event.queryStringParameters
		? event.queryStringParameters.office_name
		: null;
	console.log(rtoOfficeName);
	if (!rtoOfficeName) {
		return {
			statusCode: 200,
			body: JSON.stringify({ message: "Connected successfully" }),
		};
	}

	try {
		const params = {
			TableName: process.env.USER_TABLE_ARN,
			FilterExpression: "office_name = :officeName",
			ExpressionAttributeValues: {
				":officeName": rtoOfficeName,
			},
		};

		const data = await ddb.scan(params).promise();

		if (data.Items.length === 0) {
			return {
				statusCode: 404,
				body: JSON.stringify({
					message: `Office name '${rtoOfficeName}' is not registered`,
				}),
			};
		}
	} catch (err) {
		console.error("Error checking office name:", err);
		return {
			statusCode: 500,
			body: JSON.stringify({ error: "Error checking office name" }),
		};
	}

	try {
		console.log("In Connect: " + rtoOfficeName);
		await ddb
			.put({
				TableName: process.env.RTO_OFFICE_TABLE_ARN,
				Item: {
					office_name: rtoOfficeName,
					connectionId: event.requestContext.connectionId,
				},
			})
			.promise();
		return {
			statusCode: 200,
			body: JSON.stringify({ message: "Connected successfully" }),
		};
	} catch (err) {
		console.error("Error connecting:", err);
		return {
			statusCode: 500,
			body: JSON.stringify({ error: "Error connecting" }),
		};
	}
};
