#ifndef IPIFY_H
#define IPIFY_H

#include "IpProvider.h"

class IPIFY : public IpProvider
{
public:
    std::string url() const override {
        return "https://api64.ipify.org";
    }

    std::string method() const override {
        return "GET";
    }

    std::uint8_t ipv() const override {
        return 6;
    }

    std::string parse(std::string response) {
        return response;
    }
};

#endif // IPIFY_H
