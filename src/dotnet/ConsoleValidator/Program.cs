using ProtoValidate.Messages;
using System;
using Google.Protobuf.Reflection;

// See https://aka.ms/new-console-template for more information
class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Hello, World!");

        FileDescriptor fileDescriptor = User.Descriptor.File;

        // Print the file name
        Console.WriteLine("File Name: " + fileDescriptor.Name);

        // Iterate over all message types in the file
        foreach (var messageType in fileDescriptor.MessageTypes)
        {
            Console.WriteLine("Message Type: " + messageType.Name);

            // Iterate over all fields in the message type
            foreach (var field in messageType.Fields.InDeclarationOrder())
            {
                Console.WriteLine($"  Field: {field.Name} (Type: {field.FieldType})");
            }
        }

        // https://cloud.google.com/dotnet/docs/reference/Google.Protobuf/latest/Google.Protobuf.Reflection.FileDescriptorProto#Google_Protobuf_Reflection_FileDescriptorProto_Descriptor
        // // Access a specific message descriptor
        // Descriptor messageDescriptor = User.Descriptor;

        // // Print the message name
        // Console.WriteLine("Message Name: " + messageDescriptor.Name);

        // // Iterate over all fields in the message
        // foreach (var field in messageDescriptor.Fields.InDeclarationOrder())
        // {
        //     Console.WriteLine($"  Field: {field.Name} (Type: {field.FieldType})");
        // }
    }
}
