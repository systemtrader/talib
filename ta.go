package talib

// TA-LIB api calls

//sys ta_Initialize() (retCode int, err error) = talib.Z_Initialize
//sys ta_Shutdown() (retCode int, err error) = talib.Z_Shutdown
//sys ta_Acos(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Acos 
//sys ta_Ad(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, inVolume *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ad 
//sys ta_Add(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Add 
//sys ta_AdOsc(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, inVolume *float64, optInFastPeriod int, optInSlowPeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_AdOsc 
//sys ta_Adx(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Adx 
//sys ta_Adxr(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Adxr 
//sys ta_Apo(startIdx int, endIdx int, inReal *float64, optInFastPeriod int, optInSlowPeriod int, optInMAType int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Apo 
//sys ta_AroOn(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outAroonDown *float64, outAroonUp *float64) (retCode int, err error) = talib.Z_AroOn 
//sys ta_AroOnOsc(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_AroOnOsc 
//sys ta_AsIn(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_AsIn 
//sys ta_Atan(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Atan 
//sys ta_Atr(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Atr 
//sys ta_AvgPrice(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_AvgPrice 
//sys ta_BBands(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInNbDevUp float64, optInNbDevDn float64, optInMAType int, outBegIdx *int, outNBElement *int, outRealUpperBand *float64, outRealMiddleBand *float64, outRealLowerBand *float64) (retCode int, err error) = talib.Z_BBands 
//sys ta_Beta(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Beta 
//sys ta_Bop(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Bop 
//sys ta_Cci(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Cci 
//sys ta_Cdl2Crows(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl2Crows 
//sys ta_Cdl3BlackCrows(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3BlackCrows 
//sys ta_Cdl3InSide(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3InSide 
//sys ta_Cdl3LIneStrike(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3LIneStrike 
//sys ta_Cdl3OutSide(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3OutSide 
//sys ta_Cdl3StarsInSouth(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3StarsInSouth 
//sys ta_Cdl3WhiteSoldiers(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_Cdl3WhiteSoldiers 
//sys ta_CdlAbandOnedBaBy(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlAbandOnedBaBy 
//sys ta_CdlAdvanceBlock(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlAdvanceBlock 
//sys ta_CdlBeltHold(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlBeltHold 
//sys ta_CdlBreakaway(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlBreakaway 
//sys ta_CdlclosIngMarubozu(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlclosIngMarubozu 
//sys ta_CdlCOncealBaBySwall(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlCOncealBaBySwall 
//sys ta_CdlCounterattack(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlCounterattack 
//sys ta_CdlDarkCloudCover(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlDarkCloudCover 
//sys ta_CdlDoji(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlDoji 
//sys ta_CdlDojiStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlDojiStar 
//sys ta_CdlDragOnflyDoji(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlDragOnflyDoji 
//sys ta_CdlengulfIng(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlengulfIng 
//sys ta_CdlevenIngDojiStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlevenIngDojiStar 
//sys ta_CdlevenIngStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlevenIngStar 
//sys ta_CdlGapSidesideWhite(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlGapSidesideWhite 
//sys ta_CdlGravestOneDoji(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlGravestOneDoji 
//sys ta_CdlHammer(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHammer 
//sys ta_CdlhangIngMan(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlhangIngMan 
//sys ta_CdlHarami(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHarami 
//sys ta_CdlHaramiCross(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHaramiCross 
//sys ta_CdlHighWave(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHighWave 
//sys ta_CdlHikkake(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHikkake 
//sys ta_CdlHikkakeMod(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlHikkakeMod 
//sys ta_CdlhomIngPigeOn(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlhomIngPigeOn 
//sys ta_CdlIdentical3Crows(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlIdentical3Crows 
//sys ta_CdlInNeck(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlInNeck 
//sys ta_CdlInvertedHammer(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlInvertedHammer 
//sys ta_CdlkickIng(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlkickIng 
//sys ta_CdlkickIngByLength(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlkickIngByLength 
//sys ta_CdlLadderBottom(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlLadderBottom 
//sys ta_CdlLOngLeggedDoji(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlLOngLeggedDoji 
//sys ta_CdlLOngLIne(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlLOngLIne 
//sys ta_CdlMarubozu(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlMarubozu 
//sys ta_CdlMatchIngLow(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlMatchIngLow 
//sys ta_CdlMatHold(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlMatHold 
//sys ta_CdlmornIngDojiStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlmornIngDojiStar 
//sys ta_CdlmornIngStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, optInPenetration float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlmornIngStar 
//sys ta_CdlOnNeck(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlOnNeck 
//sys ta_CdlpiercIng(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlpiercIng 
//sys ta_CdlRickshawMan(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlRickshawMan 
//sys ta_CdlRiseFall3Methods(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlRiseFall3Methods 
//sys ta_CdlseparatIngLines(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlseparatIngLines 
//sys ta_CdlshootIngStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlshootIngStar 
//sys ta_CdlShortLIne(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlShortLIne 
//sys ta_CdlspInningTop(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlspInningTop 
//sys ta_CdlStalledPattern(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlStalledPattern 
//sys ta_CdlStickSandwich(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlStickSandwich 
//sys ta_CdlTakuri(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlTakuri 
//sys ta_CdlTasukiGap(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlTasukiGap 
//sys ta_CdlthrustIng(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlthrustIng 
//sys ta_CdltriStar(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdltriStar 
//sys ta_CdlUnique3River(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlUnique3River 
//sys ta_CdlupSideGap2Crows(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlupSideGap2Crows 
//sys ta_CdlxSideGap3Methods(startIdx int, endIdx int, inOpen *float64, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_CdlxSideGap3Methods 
//sys ta_Ceil(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ceil 
//sys ta_Cmo(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Cmo 
//sys ta_Correl(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Correl 
//sys ta_Cos(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Cos 
//sys ta_Cosh(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Cosh 
//sys ta_Dema(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Dema 
//sys ta_Div(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Div 
//sys ta_Dx(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Dx 
//sys ta_Ema(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ema 
//sys ta_Exp(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Exp 
//sys ta_Floor(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Floor 
//sys ta_HtDcPeriod(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_HtDcPeriod 
//sys ta_HtDcPhase(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_HtDcPhase 
//sys ta_HtPhasor(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outInPhase *float64, outQuadrature *float64) (retCode int, err error) = talib.Z_HtPhasor 
//sys ta_HtSIne(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outSine *float64, outLeadSine *float64) (retCode int, err error) = talib.Z_HtSIne 
//sys ta_HtTrendLIne(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_HtTrendLIne 
//sys ta_HtTrendMode(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_HtTrendMode 
//sys ta_Kama(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Kama 
//sys ta_LInearReg(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_LInearReg 
//sys ta_LInearRegAngle(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_LInearRegAngle 
//sys ta_LInearRegIntercept(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_LInearRegIntercept 
//sys ta_LInearRegSlope(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_LInearRegSlope 
//sys ta_Ln(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ln 
//sys ta_Log10(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Log10 
//sys ta_Ma(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInMAType int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ma 
//sys ta_Macd(startIdx int, endIdx int, inReal *float64, optInFastPeriod int, optInSlowPeriod int, optInSignalPeriod int, outBegIdx *int, outNBElement *int, outMACD *float64, outMACDSignal *float64, outMACDHist *float64) (retCode int, err error) = talib.Z_Macd 
//sys ta_MacdExt(startIdx int, endIdx int, inReal *float64, optInFastPeriod int, optInFastMAType int, optInSlowPeriod int, optInSlowMAType int, optInSignalPeriod int, optInSignalMAType int, outBegIdx *int, outNBElement *int, outMACD *float64, outMACDSignal *float64, outMACDHist *float64) (retCode int, err error) = talib.Z_MacdExt 
//sys ta_MacdFix(startIdx int, endIdx int, inReal *float64, optInSignalPeriod int, outBegIdx *int, outNBElement *int, outMACD *float64, outMACDSignal *float64, outMACDHist *float64) (retCode int, err error) = talib.Z_MacdFix 
//sys ta_Mama(startIdx int, endIdx int, inReal *float64, optInFastLimit float64, optInSlowLimit float64, outBegIdx *int, outNBElement *int, outMAMA *float64, outFAMA *float64) (retCode int, err error) = talib.Z_Mama 
//sys ta_Mavp(startIdx int, endIdx int, inReal *float64, inPeriods *float64, optInMinPeriod int, optInMaxPeriod int, optInMAType int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Mavp 
//sys ta_Max(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Max 
//sys ta_MaxIndex(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_MaxIndex 
//sys ta_MedPrice(startIdx int, endIdx int, inHigh *float64, inLow *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MedPrice 
//sys ta_Mfi(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, inVolume *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Mfi 
//sys ta_MidpoInt(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MidpoInt 
//sys ta_MidPrice(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MidPrice 
//sys ta_MIn(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MIn 
//sys ta_MInIndex(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outInteger *int) (retCode int, err error) = talib.Z_MInIndex 
//sys ta_MInMax(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outMin *float64, outMax *float64) (retCode int, err error) = talib.Z_MInMax 
//sys ta_MInMaxIndex(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outMinIdx *int, outMaxIdx *int) (retCode int, err error) = talib.Z_MInMaxIndex 
//sys ta_MInusDi(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MInusDi 
//sys ta_MInusDm(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_MInusDm 
//sys ta_Mom(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Mom 
//sys ta_Mult(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Mult 
//sys ta_Natr(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Natr 
//sys ta_Obv(startIdx int, endIdx int, inReal *float64, inVolume *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Obv 
//sys ta_PlusDi(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_PlusDi 
//sys ta_PlusDm(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_PlusDm 
//sys ta_Ppo(startIdx int, endIdx int, inReal *float64, optInFastPeriod int, optInSlowPeriod int, optInMAType int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Ppo 
//sys ta_Roc(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Roc 
//sys ta_Rocp(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Rocp 
//sys ta_Rocr(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Rocr 
//sys ta_Rocr100(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Rocr100 
//sys ta_Rsi(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Rsi 
//sys ta_Sar(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInAcceleration float64, optInMaximum float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sar 
//sys ta_SarExt(startIdx int, endIdx int, inHigh *float64, inLow *float64, optInStartValue float64, optInOffsetOnReverse float64, optInAccelerationInitLong float64, optInAccelerationLong float64, optInAccelerationMaxLong float64, optInAccelerationInitShort float64, optInAccelerationShort float64, optInAccelerationMaxShort float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_SarExt 
//sys ta_SIn(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_SIn 
//sys ta_SInh(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_SInh 
//sys ta_Sma(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sma 
//sys ta_Sqrt(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sqrt 
//sys ta_StdDev(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInNbDev float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_StdDev 
//sys ta_Stoch(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInFastK_Period int, optInSlowK_Period int, optInSlowK_MAType int, optInSlowD_Period int, optInSlowD_MAType int, outBegIdx *int, outNBElement *int, outSlowK *float64, outSlowD *float64) (retCode int, err error) = talib.Z_Stoch 
//sys ta_Stochf(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInFastK_Period int, optInFastD_Period int, optInFastD_MAType int, outBegIdx *int, outNBElement *int, outFastK *float64, outFastD *float64) (retCode int, err error) = talib.Z_Stochf 
//sys ta_StochRsi(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInFastK_Period int, optInFastD_Period int, optInFastD_MAType int, outBegIdx *int, outNBElement *int, outFastK *float64, outFastD *float64) (retCode int, err error) = talib.Z_StochRsi 
//sys ta_Sub(startIdx int, endIdx int, inReal0 *float64, inReal1 *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sub 
//sys ta_Sum(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Sum 
//sys ta_T3(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInVFactor float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_T3 
//sys ta_Tan(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Tan 
//sys ta_Tanh(startIdx int, endIdx int, inReal *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Tanh 
//sys ta_Tema(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Tema 
//sys ta_Trange(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Trange 
//sys ta_Trima(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Trima 
//sys ta_Trix(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Trix 
//sys ta_Tsf(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Tsf 
//sys ta_TypPrice(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_TypPrice 
//sys ta_UltOsc(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod1 int, optInTimePeriod2 int, optInTimePeriod3 int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_UltOsc 
//sys ta_Var(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, optInNbDev float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Var 
//sys ta_WclPrice(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_WclPrice 
//sys ta_Willr(startIdx int, endIdx int, inHigh *float64, inLow *float64, inClose *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Willr 
//sys ta_Wma(startIdx int, endIdx int, inReal *float64, optInTimePeriod int, outBegIdx *int, outNBElement *int, outReal *float64) (retCode int, err error) = talib.Z_Wma 
