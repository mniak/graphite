#pragma once

#include "declarations.h"
#include <vector>

using namespace std;

class Program {
    public :
    vector<ExternalLibraryDeclaration> libraries;
    vector<IStatement> entrypoint;
};