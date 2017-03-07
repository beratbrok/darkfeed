#include <cstdint>
#include "darkfeed/util/crc32c.h"
#include "crc32ctables.h"

namespace darkfeed
{
std::uint32_t crc32cSlicingBy8(std::uint32_t crc, const void *data, size_t length)
{
    const char *p_buf = (const char *) data;

    // Handle leading misaligned bytes
    size_t initial_bytes = (sizeof(int32_t) - (intptr_t) p_buf) & (sizeof(int32_t) - 1);
    if (length < initial_bytes)
        initial_bytes = length;
    for (size_t li = 0; li < initial_bytes; li++)
    {
        crc = crc_tableil8_o32[(crc ^ *p_buf++) & 0x000000FF] ^ (crc >> 8);
    }

    length -= initial_bytes;
    size_t running_length = length & ~(sizeof(uint64_t) - 1);
    size_t end_bytes = length - running_length;

    for (size_t li = 0; li < running_length / 8; li++)
    {
        crc ^= *(uint32_t *) p_buf;
        p_buf += 4;
        uint32_t term1 = crc_tableil8_o88[crc & 0x000000FF] ^
                         crc_tableil8_o80[(crc >> 8) & 0x000000FF];
        uint32_t term2 = crc >> 16;
        crc = term1 ^
              crc_tableil8_o72[term2 & 0x000000FF] ^
              crc_tableil8_o64[(term2 >> 8) & 0x000000FF];
        term1 = crc_tableil8_o56[(*(uint32_t *) p_buf) & 0x000000FF] ^
                crc_tableil8_o48[((*(uint32_t *) p_buf) >> 8) & 0x000000FF];

        term2 = (*(uint32_t *) p_buf) >> 16;
        crc = crc ^ term1 ^
              crc_tableil8_o40[term2 & 0x000000FF] ^
              crc_tableil8_o32[(term2 >> 8) & 0x000000FF];
        p_buf += 4;
    }

    for (size_t li = 0; li < end_bytes; li++)
    {
        crc = crc_tableil8_o32[(crc ^ *p_buf++) & 0x000000FF] ^ (crc >> 8);
    }

    return crc;
}
} //darkfeed
