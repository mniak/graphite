#pragma once
#include <string>

using namespace std;

class IValue
{
};

class StringLiteral : public IValue
{
public:
    string value;
    StringLiteral(string value)
    {
        this->value = value;
    };
};