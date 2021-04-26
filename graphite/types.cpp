#include "types.h"

using namespace std;

PrimitiveType::PrimitiveType(string name)
{
    this->name = name;
}
const PrimitiveType _String("string");
PrimitiveType PrimitiveType::String()
{
    return _String;
}