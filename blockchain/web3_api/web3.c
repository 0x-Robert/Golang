#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <curl/curl.h>
#include <jansson.h>

#define INFURA_URL "https://goerli.infura.io/v3/8f0f6f1d3eba4554818d7e69ef83053f"

struct buffer {
    char *data;
    size_t size;
};

static size_t write_callback(char *ptr, size_t size, size_t nmemb, struct buffer *buffer) {
    size_t total_size = size * nmemb;
    buffer->data = realloc(buffer->data, buffer->size + total_size + 1);
    if (buffer->data == NULL) {
        fprintf(stderr, "Error: failed to allocate memory.\n");
        return 0;
    }
    memcpy(buffer->data + buffer->size, ptr, total_size);
    buffer->size += total_size;
    buffer->data[buffer->size] = '\0';
    return total_size;
}

json_t *get_transactions_by_account(const char *account, uint64_t start_block, uint64_t end_block) {
    CURL *curl_handle = NULL;
    CURLcode res;
    json_t *transactions = json_array();

    struct buffer buffer = { NULL, 0 };

    // create a CURL handle
    curl_handle = curl_easy_init();
    if (curl_handle == NULL) {
        fprintf(stderr, "Error: failed to create CURL handle.\n");
        return NULL;
    }

    // set the CURL options
    char url[512];
    sprintf(url, "%s/%s", INFURA_URL, "eth_getBlockByNumber");
    curl_easy_setopt(curl_handle, CURLOPT_URL, url);
    curl_easy_setopt(curl_handle, CURLOPT_FOLLOWLOCATION, 1L);
    curl_easy_setopt(curl_handle, CURLOPT_WRITEFUNCTION, write_callback);
    curl_easy_setopt(curl_handle, CURLOPT_WRITEDATA, &buffer);

    // loop over all the blocks in the specified range
    for (uint64_t block_number = start_block; block_number <= end_block; block_number++) {
        // create the JSON-RPC request object
        json_t *request = json_pack("{s:s, s:s, s:[{s:s, s:s}]}", "jsonrpc", "2.0", "method", "eth_getBlockByNumber", "params", "[]", "id", "1");
        json_array_append_new(json_object_get(request, "params"), json_string_printf("0x%" PRIx64, block_number));
        json_array_append_new(json_object_get(request, "params"), json_true());

        // encode the JSON-RPC request object as a string
        char *request_string = json_dumps(request, JSON_COMPACT);
        size_t request_string_length = strlen(request_string);

        // set the request headers
        struct curl_slist *headers = NULL;
        headers = curl_slist_append(headers, "Content-Type: application/json");

        // set the request body
        curl_easy_setopt(curl_handle, CURLOPT_POST, 1L);
        curl_easy_setopt(curl_handle, CURLOPT_POSTFIELDS, request_string);
        curl_easy_setopt(curl_handle, CURLOPT_POSTFIELDSIZE, (long) request_string_length);

        // send the HTTP request
        res = curl_easy_perform(curl_handle);
        if (res != CURLE_OK) {
            fprintf(stderr, "Error
        }
        }
    