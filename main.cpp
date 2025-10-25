#include <iostream>
#include <string>
#include <sstream>
#include <fstream>

#include "json.hpp"
#include "LSPTypes.h"

using json = nlohmann::json;

void log(std::string message){
    std::ofstream logFile;
    
    logFile.open("logs.txt", std::ios_base::app);
    logFile << message << "\n\n";
    logFile.close();
}

std::string readMessage(){
    std::string line;
    int contentLength = -1;

    while(std::getline(std::cin, line)){

        if(!line.empty() && line.back() == '\r'){
            line.pop_back();
        }

        if(line.empty()){
            break;
        }

        if(line.find("Content-Length:") == 0){
            std::string lengthStr = line.substr(15);
            
            size_t start = lengthStr.find_first_not_of(" \t");
            if(start != std::string::npos){
                lengthStr = lengthStr.substr(start);
            }

            try{
                contentLength = std::stoi(lengthStr);
            
            }catch (const std::exception& e) {
                log("[ERR] Invalid Content-Length value: " + lengthStr);
                
                return "";
            }
        }
    }

    if(contentLength <= 0){
        log("[ERR] Invalid or missing Content-Length");
        
        return "";
    }

    std::string body(contentLength, '\0');
    std::cin.read(&body[0], contentLength);

    if(std::cin.gcount() != contentLength){
        log("[ERR] Did not read expected content length. Got: " + std::to_string(std::cin.gcount()));
        
        return "";
    }

    return body;
}

void sendMessage(json msg){
    std::string body = msg.dump();
    
    std::cout << "Content-Length: " << body.size() << "\r\n\r\n" << body;

    std::cout.flush();
}

int main(){
    while(true){
        std::string clientMessage = readMessage();

        if(clientMessage.empty()){
            continue;
        }

        json clientMessageJSON = json::parse(clientMessage);

        if(clientMessageJSON.is_discarded()){
            continue;
        }

        // TODO: Abstract all this and stuffs
        if(clientMessageJSON.contains("method")){
            std::string method = clientMessageJSON["method"];

            if(method == "initialize"){
                log("INITIALISED SUCCESSFULLY");

                InitialiseResponse response(clientMessageJSON["id"]);

                sendMessage(response.toJson());
            }
        }
    }

    return 0;
}