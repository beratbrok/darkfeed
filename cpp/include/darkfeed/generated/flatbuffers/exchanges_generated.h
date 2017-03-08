// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_EXCHANGES_DARKFEED_SCHEMAS_FB_H_
#define FLATBUFFERS_GENERATED_EXCHANGES_DARKFEED_SCHEMAS_FB_H_

#include "flatbuffers/flatbuffers.h"

namespace darkfeed
{
namespace schemas
{
namespace fb
{

/// @brief MIC Codes for exchanges
enum class MIC : uint8_t {
    ///< No Exchange
        NONE = 0  ///< NASDAQ (Used for displayed trades/quotes)
    ,
    XNAS = 1  ///< New York Stock Exchange
    ,
    XNYS = 2  ///< NYSE Alternext (Formerly American Stock Exchange)
    ,
    XASE = 3  ///< Chicago Board of Options Exchange
    ,
    CBOE = 4  ///< NYSE ARCA (Formerly ArcaEx, Archipelago Exchange)
    ,
    ARCX = 5  ///< National Stock Exchange (Formerly Cincinatti Stock Exchange)
    ,
    XCIS = 6  ///< NASDAQ OMX PHLX (Formerly Philadelphia Stock Exchange)
    ,
    XPHL = 7  ///< Options Pricing Reporting Authority
    ,
    OPRA = 8  ///< NASDAQ OMX BX (Formerly BSE, Boston Stock Exchange)
    ,
    XBOS = 9  ///< NASDAQ Global/Select (Listing MIC for securities meeting all of NASDAQ's listing requirements. Quotes/trades reported by XNAS)
    ,
    XNGS = 10  ///< NASDAQ Capital Market (Listing MIC for small-cap securities meeting streamlined requirements. Quotes/trades reported by XNAS)
    ,
    XNCM = 11  ///< BATS Trading
    ,
    BATS = 12  ///< BATS Trading Y-Exchange
    ,
    BATY = 13  ///< Direct Edge Exchange (Operated by BATS Trading)
    ,
    EDGE = 14  ///< Direct Edge X Exchange (Operated by BATS Trading)
    ,
    EDGX = 15  ///< Chicago Stock Exchange
    ,
    XCHI = 16  ///< OTC Bulletin Board (Operated by NASDAQ)
    ,
    XOTC = 17  ///< Direct Edge A Exchange
    ,
    EDGA = 18  ///< SMART Order Router (Non-MIC)
    ,
    SMART = 19  ///< ISLAND ECN
    ,
    ICEL = 20  ///< Investors Exchange (IEX)
    ,
    IEX = 21  ///< SIP (Standard Information Processor)
    ,
    SIP = 22,
    MIN = NONE,
    MAX = SIP
};

inline const char **EnumNamesMIC()
{
  static const char *names[] = {
      "NONE",
      "XNAS",
      "XNYS",
      "XASE",
      "CBOE",
      "ARCX",
      "XCIS",
      "XPHL",
      "OPRA",
      "XBOS",
      "XNGS",
      "XNCM",
      "BATS",
      "BATY",
      "EDGE",
      "EDGX",
      "XCHI",
      "XOTC",
      "EDGA",
      "SMART",
      "ICEL",
      "IEX",
      "SIP",
      nullptr
  };
  return names;
}

inline const char *EnumNameMIC(MIC e)
{
  const size_t index = static_cast<int>(e);
  return EnumNamesMIC()[index];
}

}  // namespace fb
}  // namespace schemas
}  // namespace darkfeed

#endif  // FLATBUFFERS_GENERATED_EXCHANGES_DARKFEED_SCHEMAS_FB_H_