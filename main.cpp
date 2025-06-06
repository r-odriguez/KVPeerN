#include <QApplication>
#include <map>
#include <string>

#include "MainWindow.h"

class ip_provider
{
private:
	std::string url;
	std::string method;
	int ipv;
	std::map<std::string, std::string> opt;
	ip_provider_response res;

public:
	std::string ip;
};

// IPIFY ----------------------------------------------------------------------
class IPIFY: ip_provider
{
private:
	std::string url = "https://api64.ipify.org";
	std::string method = "GET";
	int ipv = 6;
	std::map<std::string, std::string> opt;
	ip_provider_response res;

public:
	virtual std::string parse(std::string response)
	{
		this->ip = response;
		return response;
	}
};

// IFCFG ----------------------------------------------------------------------
class IFCFG: ip_provider
{
private:
	std::string url = "https://ifcfg.me/";
	std::string method = "GET";
	int ipv = 4;
	std::map<std::string, std::string> opt;
	ip_provider_response res;

public:
	virtual std::string parse(std::string response)
	{
		this->ip = response;
		return response;
	}
}

int main(int argc, char *argv[])
{
	std::string PUBLIC_IP_PROVIDERS[2] = { "ifcfg", "ipify" };

	QApplication a(argc, argv);
	MainWindow w;
	w.show();
	return a.exec();
}
