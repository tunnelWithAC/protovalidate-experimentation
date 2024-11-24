const FooRequest = require("./proto/foo_pb");
const { validate } = require("protovalidate");

const rules = {
 field: "bar",
 type: "string",
 required: true,
};

const schema = {
 FooRequest: [rules],
};

const validator = validate(schema);

const fooRequest = new FooRequest();
fooRequest.setBar("Hello, World!");
fooRequest.setBaz(42);

const validationResult = validator.validate(fooRequest);

if (validationResult.isValid) {
 console.log("FooRequest is valid.");
} else {
 console.log("FooRequest is invalid:", validationResult.errors);
}

