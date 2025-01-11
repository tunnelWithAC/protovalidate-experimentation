using DotProtoValidate.Messages;
using ConsoleValidator.Validators;

namespace DotProtoValidate.Test;

public class ProtoValidatorTests
{
    private ProtoValidator _validator;


    [SetUp]
    public void Setup()
    {
         _validator = new ProtoValidator();

    }

    [Test]
    public void ValidateUser_ValidUser_ReturnsSuccess()
    {
        var user = new User { Age = 18 };

        var result = _validator.ValidateUser(user);

        Assert.IsTrue(result.IsSuccess);
    }

    [Test]
    public void ValidateUser_InvalidUser_ReturnsError()
    {
        var user = new User { Age = 17 };

        var result = _validator.ValidateUser(user);

        Assert.IsFalse(result.IsSuccess);
    }
}
