all: avsc
avsc:
	mkdir avro && gogen-avro avro/ blobschema.avsc 
