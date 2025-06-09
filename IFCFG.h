#ifndef IFCFG_H
#define IFCFG_H

#include "IpProvider.h"

class IFCFG : public IpProvider
{
public:
    std::string url() const override {
        return "https://ifcfg.me/";
    }

    std::string method() const override {
        return "GET";
    }

    std::uint8_t ipv() const override {
        return 4;
    }

    std::string parse(std::string response) {
        return response;
    }
};

#endif // IFCFG_H
