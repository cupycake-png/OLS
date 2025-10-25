#pragma once

#include <string>

#include "json.hpp"

class Message{
    protected:
        std::string jsonrpc = "2.0";

    public:
        virtual ~Message() = default;
        virtual nlohmann::json toJson() = 0;
};

class RequestMessage : public Message{
    protected:
        int id = 0;
        std::string method;
        nlohmann::json params;

    public:
        RequestMessage(std::string method, int id=0) : method(method), id(id) {}

        nlohmann::json toJson() override{
            return {
                {"jsonrpc", jsonrpc},
                {"id", id},
                {"method", method},
                {"params", params}
            };
        }

        void setParams(nlohmann::json& p){
            params = p;
        }
};

class ResponseMessage : public Message{
    protected:
        int id = 0;
        nlohmann::json result;
        nlohmann::json error;

    public:
        ResponseMessage(int id=0) : id(id) {}

        void setResult(nlohmann::json& r){
            result = r;
        }

        void setError(nlohmann::json& e){
            error = e;
        }

        nlohmann::json toJson() override{
            nlohmann::json ret = {
                {"jsonrpc", jsonrpc},
                {"id", id}
            };

            if(!result.is_null()){
                ret["result"] = result;
            }

            if(!error.is_null()){
                ret["error"] = error;
            }

            return ret;
        }
};

class InitialiseRequest : public RequestMessage{
    public:
        InitialiseRequest(int id) : RequestMessage("initialize", id){
            params = {
                {"processId", NULL},
                {"clientInfo", {
                    {"name", "cupycode"},
                    {"version", "1.0"}
                }},
                {"locale", "en"}
            };
        }
};

class InitialiseResponse : public ResponseMessage{
    public:
        InitialiseResponse(int id) : ResponseMessage(id){
            result = {
                {"serverInfo", {
                    {"name", "OLS"},
                    {"version", "1.0"}
                }}
            };
        }
};