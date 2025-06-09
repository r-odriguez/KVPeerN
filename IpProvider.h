#include <string>
#include <cstdint>
#include <map>

#ifndef IPPROVIDER_H
#define IPPROVIDER_H

class IpProvider
{
public:
    virtual ~IpProvider() = default;

    virtual std::string url() const = 0;
    virtual std::string method() const = 0;
    virtual std::uint8_t ipv() const = 0;
    virtual std::map<std::string, std::string> options() const = 0;

    virtual std::string parse(std::string response) const = 0;
};

#endif // IPPROVIDER_H
