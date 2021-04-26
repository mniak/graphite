#pragma once

#include <string>

using namespace std;

class Type
{
};

class PrimitiveType : public Type
{
public:
    string name;
    static PrimitiveType String();
    PrimitiveType(string name);
};
