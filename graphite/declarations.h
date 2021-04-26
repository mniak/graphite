#pragma once

#include "types.h"
#include <string>
#include <vector>

using namespace std;

class IMethodDeclaration
{
public:
    string name;
};

class MethodParameterDeclaration
{
public:
    string name;
    Type type;
};

class IStatement
{
};

class MethodDeclaration : IMethodDeclaration
{
public:
    // string name;
    vector<MethodParameterDeclaration> parameters;
    vector<IStatement> statements;
};

class ModuleDeclaration
{
public:
    string name;
    vector<MethodDeclaration> methods;
};

class LibraryDeclaration
{
public:
    string name;
    vector<ModuleDeclaration> modules;
};

class ExternalMethodDeclaration : public IMethodDeclaration
{
public:
    // string name;
    vector<MethodParameterDeclaration> parameters;
};

class ExternalModuleDeclaration
{
public:
    string name;
    vector<ExternalMethodDeclaration> methods;
};


class ExternalLibraryDeclaration
{
public:
    string name;
    vector<ExternalModuleDeclaration> modules;
};