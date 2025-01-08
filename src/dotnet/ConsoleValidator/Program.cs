﻿
using ConsoleValidator.Messages;
using System;
using Google.Protobuf.Reflection;
using ProtoValidate;

// See https://aka.ms/new-console-template for more information
class Program
{
    static void Main(string[] args)
    {
        var validatorOptions = new ProtoValidate.ValidatorOptions() {
            // This setting is used to configure if it loads your validation descriptors upon creation of the validator.
            // True will load on creation
            // False will defer loading the validator until first run of the validation logic for that type.
            PreLoadDescriptors = false,

            // This setting will cause a compilation exception to be thrown if the message type you are validating hasn't been pre-loaded using the file descriptor list.
            DisableLazy = false,

            //register your file descriptors generated by Google.Protobuf library for your compiled .proto files
            FileDescriptors = new List<FileDescriptor>() {
                //your list of Protobuf File Descriptors here
                User.Descriptor.File
            }
        };

        //Instantiate the validator.  You should cache the validator for reuse.
        var validator = new ProtoValidate.Validator(validatorOptions);
        Console.WriteLine(validator.GetType().Name);

        // flag to indicate if the validator should return on the first error (true) or validate all the fields and return all the errors in the message (false).
        var failFast = false;

        var user = new User { Age = 17 };

        // //validate the message
        var violations = validator.Validate(user, failFast);

        // //the violations contains the validation errors.
        var hasViolations = violations.Violations.Count > 0;
        Console.WriteLine(hasViolations.ToString());

        foreach(var violation in violations.Violations)
        {
            Console.WriteLine("Violation: {0}", violation);
        }
    }
}
