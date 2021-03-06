#include "darkfeed/util/hashing.hpp"
#include <gtest/gtest.h>

using namespace darkfeed;


TEST(CRC32C, SliceBy8Empty)
{
    const char *empty = "";
    unsigned char expected_arr[] = {0x00, 0x00, 0x00, 0x00};
    std::uint32_t crc = CRC32C_FINISH(crc32cSlicingBy8(CRC32C_INIT, empty, 0));
    ASSERT_EQ(*(std::uint32_t *) expected_arr, crc);
}


TEST(CRC32C, SliceBy8Alphabet)
{
    const char alphabet[] = "abcdefghijklmnopqrstuvwxyz";
    unsigned char expected_arr[] = {0x25, 0xef, 0xe6, 0x9e};
    std::uint32_t crc = CRC32C_FINISH(crc32cSlicingBy8(CRC32C_INIT, alphabet, 26));
    ASSERT_EQ(*(std::uint32_t *) expected_arr, crc);
}

#ifdef HAS_SSE_4_2


TEST(CRC32C, IntelEmpty)
{
    const char *empty = "";
    unsigned char expected_arr[] = {0x00, 0x00, 0x00, 0x00};
    std::uint32_t crc = CRC32C_FINISH(crc32cIntelC(CRC32C_INIT, empty, 0));
    ASSERT_EQ(*(std::uint32_t *) expected_arr, crc);
}


TEST(CRC32C, IntelAlphabet)
{
    const char alphabet[] = "abcdefghijklmnopqrstuvwxyz";
    unsigned char expected_arr[] = {0x25, 0xef, 0xe6, 0x9e};
    std::uint32_t crc = CRC32C_FINISH(crc32cIntelC(CRC32C_INIT, alphabet, 26));
    ASSERT_EQ(*(std::uint32_t *) expected_arr, crc);
}
#endif


TEST(CRC32CCStrHasher, Alphabet)
{
    const char alphabet[] = "abcdefghijklmnopqrstuvwxyz";
    unsigned char expected_arr[] = {0x25, 0xef, 0xe6, 0x9e};
    CRC32CCStrHasher hasher;
    std::uint32_t crc = hasher(alphabet);
    ASSERT_EQ(*(std::uint32_t *) expected_arr, crc);
}


TEST(CRC32CStrHasher, Alphabet)
{
    std::string alphabet = "abcdefghijklmnopqrstuvwxyz";
    unsigned char expected_arr[] = {0x25, 0xef, 0xe6, 0x9e};
    CRC32CStrHasher hasher;
    std::uint32_t crc = hasher(alphabet);
    ASSERT_EQ(*(std::uint32_t *) expected_arr, crc);
}


TEST(CRC32CSymHasher, BaseCase)
{
    Symbol s(MIC::XNYS, "BRK.PR.A");
    ASSERT_STREQ("BRK", s.root);
    ASSERT_EQ(IssueType::preferred, s.issue_type);
    ASSERT_EQ('A', s.series);
    CRC32CSymHasher h;
    auto h1 = h(s);
    auto sc = s;
    auto h2 = h(sc);
    ASSERT_EQ(h1, h2);
    Symbol s2(MIC::XNYS, "BRK.A");
    ASSERT_STREQ("BRK", s2.root);
    ASSERT_EQ(IssueType::normal, s2.issue_type);
    ASSERT_EQ('A', s2.series);
    auto h3 = h(s2);
    ASSERT_NE(h1, h3);
}
