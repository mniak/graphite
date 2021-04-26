#pragma once

#include "declarations.h"
#include "values.h"

using namespace std;

class Argument
{
public:
    MethodParameterDeclaration parameter;
    IValue value;
    
    Argument(MethodParameterDeclaration parameter, IValue value)
    {
        this->parameter = parameter;
        this->value = value;
    };
};

class MethodInvocation : public IStatement
{
public:
    IMethodDeclaration method;
    vector<Argument> arguments;
};