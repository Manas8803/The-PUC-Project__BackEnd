const AWS = require("aws-sdk");
const ddb = new AWS.DynamoDB.DocumentClient();

exports.handler = async function (event, context) {
	const connectionId = event.requestContext.connectionId;
	try {
		console.log("In Disconnect: " + connectionId);

		// First, scan the table to find the item with the desired connectionId
		const data = await ddb
			.scan({
				TableName: process.env.RTO_OFFICE_TABLE_ARN,
				FilterExpression: "#connectionId = :connectionId",
				ExpressionAttributeNames: {
					"#connectionId": "connectionId",
				},
				ExpressionAttributeValues: {
					":connectionId": connectionId,
				},
			})
			.promise();

		if (data.Items.length === 0) {
			return {
				statusCode: 404,
				body: JSON.stringify({
					message: "No item found with the given connectionId",
				}),
			};
		}

		const itemToDelete = data.Items[0];

		await ddb
			.delete({
				TableName: process.env.RTO_OFFICE_TABLE_ARN,
				Key: {
					office_name: itemToDelete.office_name,
				},
			})
			.promise();

		return {
			statusCode: 200,
			body: JSON.stringify({ message: "Disconnected successfully" }),
		};
	} catch (err) {
		console.error("Error disconnecting:", err);
		return {
			statusCode: 500,
			body: JSON.stringify({ error: "Error disconnecting" }),
		};
	}
};
