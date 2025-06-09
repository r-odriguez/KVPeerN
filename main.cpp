#include <QApplication>
#include <vector>
#include <string>

#include "MainWindow.h"
#include "IpProvider.h"
#include "ifcfg.h"
#include "ipify.h"

int main(int argc, char *argv[])
{
    std::vector<std::unique_ptr<IpProvider>> PUBLIC_IP_PROVIDERS;
    PUBLIC_IP_PROVIDERS.push_back(std::make_unique<IFCFG>());
    PUBLIC_IP_PROVIDERS.push_back(std::make_unique<IPIFY>());

	QApplication a(argc, argv);
	MainWindow w;
	w.show();
	return a.exec();
}
