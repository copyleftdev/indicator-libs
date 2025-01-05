# Indicator Libs

A **Go (Golang) library** of popular technical analysis indicators. This is a **private** repository, providing a range of momentum and trend-based tools that can be integrated into quantitative trading or analysis workflows.

## Table of Contents

1. [SMA (Simple Moving Average)](#1-sma-simple-moving-average)  
2. [EMA (Exponential Moving Average)](#2-ema-exponential-moving-average)  
3. [RSI (Relative Strength Index)](#3-rsi-relative-strength-index)  
4. [MACD (Moving Average ConvergenceDivergence)](#4-macd-moving-average-convergencedivergence)  
5. [Bollinger Bands](#5-bollinger-bands)  
6. [Stochastic Oscillator](#6-stochastic-oscillator)  
7. [ATR (Average True Range)](#7-atr-average-true-range)  
8. [ADX (Average Directional Index)](#8-adx-average-directional-index)  
9. [CCI (Commodity Channel Index)](#9-cci-commodity-channel-index)  
10. [Williams %R](#10-williams-r)  
11. [OBV (On-Balance Volume)](#11-obv-on-balance-volume)  
12. [MFI (Money Flow Index)](#12-mfi-money-flow-index)  
13. [Ultimate Oscillator (UO)](#13-ultimate-oscillator-uo)  
14. [Ichimoku Kinko Hyo (Ichimoku Cloud)](#14-ichimoku-kinko-hyo-ichimoku-cloud)  
15. [Parabolic SAR](#15-parabolic-sar)  
16. [Keltner Channels](#16-keltner-channels)  
17. [KAMA (Kaufman’s Adaptive Moving Average)](#17-kama-kaufmans-adaptive-moving-average)  
18. [SuperTrend](#18-supertrend)  
19. [T3 (Tillson’s T3 Moving Average)](#19-t3-tillsons-t3-moving-average)

---

## 1. SMA (Simple Moving Average)

- **Origin**: A foundational smoothing technique in time-series analysis, used for decades in finance.  
- **Description**: Takes the arithmetic mean of prices over a fixed window.  
- **Common Parameters**:  
  - `window` (e.g., 20).  
- **Use Cases & Patterns**:  
  - **Trend identification** by smoothing data over time.  
  - **Crossover strategies** when used with multiple SMAs.

---

## 2. EMA (Exponential Moving Average)

- **Origin**: A refinement of SMA, giving more weight to recent data points.  
- **Description**: Reduces lag by exponentially weighting recent prices.  
- **Common Parameters**:  
  - `window` (e.g., 20).  
- **Use Cases & Patterns**:  
  - **Faster reaction** to sudden price changes.  
  - Used in **MACD** calculations and multi-EMA cross strategies.

---

## 3. RSI (Relative Strength Index)

- **Origin**: Developed by J. Welles Wilder Jr. in the late 1970s.  
- **Description**: Ranges from 0 to 100, identifying overbought (>70) or oversold (<30) market conditions.  
- **Common Parameters**:  
  - `window` (e.g., 14).  
- **Use Cases & Patterns**:  
  - Spot possible **reversals**; look for RSI crossing key thresholds.  
  - **Divergence** between RSI and price can indicate momentum shifts.

---

## 4. MACD (Moving Average Convergence/Divergence)

- **Origin**: Created by Gerald Appel (1970s).  
- **Description**: Uses two EMAs (fast & slow) and a signal line to identify momentum and potential crossovers.  
- **Common Parameters**:  
  - `fastPeriod` (e.g., 12), `slowPeriod` (e.g., 26), `signalPeriod` (9).  
- **Use Cases & Patterns**:  
  - **Crossover** signals between MACD line & signal line.  
  - **Histogram** expansions to judge trend strength.

---

## 5. Bollinger Bands

- **Origin**: Developed by John Bollinger in the early 1980s.  
- **Description**: Plots an SMA-based middle band plus upper/lower bands at a specified number of standard deviations.  
- **Common Parameters**:  
  - `window` (e.g., 20), `num_std` (commonly 2).  
- **Use Cases & Patterns**:  
  - **Volatility** assessment (width of bands).  
  - **Bollinger Squeeze** signals potential breakouts.

---

## 6. Stochastic Oscillator

- **Origin**: Created by George C. Lane in the 1950s.  
- **Description**: Compares current close to the recent range, producing %K and %D lines typically in [0,100].  
- **Common Parameters**:  
  - `k_period` (14), `d_period` (3).  
- **Use Cases & Patterns**:  
  - Identifying **overbought/oversold** conditions.  
  - **%K–%D crossovers** for entry/exit signals.

---

## 7. ATR (Average True Range)

- **Origin**: By J. Welles Wilder Jr., also in the late 1970s.  
- **Description**: Measures volatility by considering the full price range and gaps.  
- **Common Parameters**:  
  - `window` (e.g., 14).  
- **Use Cases & Patterns**:  
  - Setting stop losses or position sizing based on **market volatility**.  
  - Filtering out low-volatility periods.

---

## 8. ADX (Average Directional Index)

- **Origin**: Also by Wilder in the late 1970s.  
- **Description**: Measures trend strength on a 0–100 scale, often paired with +DI and -DI.  
- **Common Parameters**:  
  - `window` (e.g., 14).  
- **Use Cases & Patterns**:  
  - Distinguishing **strong trending** markets (ADX > 25).  
  - +DI/-DI **crossovers** for bullish/bearish signals.

---

## 9. CCI (Commodity Channel Index)

- **Origin**: Invented by Donald Lambert (1980).  
- **Description**: Shows how far the current price is from its “average” over time; can move above +100 or below -100.  
- **Common Parameters**:  
  - `window` (often 14 or 20).  
- **Use Cases & Patterns**:  
  - Overbought/oversold detection outside ±100 range.  
  - **Divergence** signals potential trend shifts.

---

## 10. Williams %R

- **Origin**: Created by Larry Williams.  
- **Description**: Similar to Stochastic, ranges from 0 to -100, indicating the close’s position relative to recent highs/lows.  
- **Common Parameters**:  
  - `window` (e.g., 14).  
- **Use Cases & Patterns**:  
  - Short-term **overbought/oversold** triggers.  
  - Often used to confirm momentum changes.

---

## 11. OBV (On-Balance Volume)

- **Origin**: Joseph Granville (1960s).  
- **Description**: Cumulative running total that adds volume on up days, subtracts on down days.  
- **Common Parameters**:  
  - Uses daily volume and close price movement.  
- **Use Cases & Patterns**:  
  - **Volume-based divergences** (OBV diverges from price).  
  - Identifying underlying strength/weakness of a trend.

---

## 12. MFI (Money Flow Index)

- **Origin**: Created by Gene Quong and Avrum Soudack; sometimes called a “volume-weighted RSI.”  
- **Description**: Ranges 0–100, factoring in both price and volume to detect overbought/oversold conditions.  
- **Common Parameters**:  
  - `window` (often 14).  
- **Use Cases & Patterns**:  
  - More sensitive than RSI (due to volume).  
  - Divergences often signal turning points.

---

## 13. Ultimate Oscillator (UO)

- **Origin**: By Larry Williams in 1976.  
- **Description**: Combines short-, medium-, and long-term price action into one oscillator (0–100).  
- **Common Parameters**:  
  - Typically uses periods 7, 14, 28 with weighting (4,2,1).  
- **Use Cases & Patterns**:  
  - Attempts to reduce false divergences by analyzing multiple timeframes.  
  - Overbought near 70–80, oversold near 20–30.

---

## 14. Ichimoku Kinko Hyo (Ichimoku Cloud)

- **Origin**: Developed by Goichi Hosoda (published in 1969).  
- **Description**: A comprehensive indicator that plots five lines (Tenkan-sen, Kijun-sen, Senkou Span A, Senkou Span B, and Chikou Span) to show momentum, potential support/resistance, and trend direction.  
- **Common Parameters**:  
  - Tenkan-sen = 9, Kijun-sen = 26, Senkou Span B = 52, Shift = 26 (typical defaults).  
- **Use Cases & Patterns**:  
  - **Cloud** (Senkou Spans A & B) as support/resistance; bullish if price is above the cloud.  
  - **Tenkan–Kijun Cross** can signal short-term momentum changes.  
  - **Chikou Span** lags, confirming trend if it’s above/below price.

---

## 15. Parabolic SAR

- **Origin**: Developed by J. Welles Wilder Jr. in the late 1970s.  
- **Description**: *Parabolic Stop and Reverse* (SAR) uses a “parabola” trailing price to highlight potential stop-loss levels and trend reversals.  
- **Common Parameters**:  
  - `accelerationFactor` (AF) starting value, often 0.02;  
  - `accelerationMax` (e.g., 0.2).  
- **Use Cases & Patterns**:  
  - **Trend-following** with automated trailing stops.  
  - Dots appear below price in an uptrend and above price in a downtrend; reversal triggers when price crosses the SAR level.

---

## 16. Keltner Channels

- **Origin**: Based on work by Chester Keltner in the 1960s, later modified/popularized by Linda Bradford Raschke.  
- **Description**: A volatility-based envelope indicator. The middle line is typically an EMA of typical price (High+Low+Close/3), and the upper/lower lines are offset by a multiple of ATR.  
- **Common Parameters**:  
  - `emaPeriod` for the middle line (e.g., 20).  
  - `atrPeriod` (e.g., 10).  
  - `mult` as a multiplier for the ATR offset (commonly around 2.0).  
- **Use Cases & Patterns**:  
  - **Volatility-based bands** that contract/expand with market moves.  
  - Similar to Bollinger Bands but uses ATR instead of standard deviation.  
  - Can signal breakouts when price strongly pierces the upper or lower channel.

---

## 17. KAMA (Kaufman’s Adaptive Moving Average)

- **Origin**: Created by Perry J. Kaufman, introduced in his 1998 book “Trading Systems and Methods” (though developed earlier).  
- **Description**: Adjusts the moving average's sensitivity based on market “noise” (volatility) vs. direction, resulting in less lag in trending markets and smoother lines during sideways action.  
- **Common Parameters**:  
  - `ERPeriod` for the efficiency ratio (e.g., 10).  
  - `FastPeriod` (e.g., 2) and `SlowPeriod` (e.g., 30) for calculating adaptive smoothing constants.  
- **Use Cases & Patterns**:  
  - **Adaptive smoothing** that reduces whipsaws in choppy markets.  
  - Reacts more quickly in strong trends, more slowly in sideways conditions.

---

## 18. SuperTrend

- **Origin**: A more recent innovation, popularized in the 2000s by various traders and widely used in algorithmic strategies.  
- **Description**: Uses **ATR** to generate upper and lower trailing stops (“bands”). The indicator flips between uptrend and downtrend when price crosses these bands.  
- **Common Parameters**:  
  - `period` for ATR (e.g., 10 or 14).  
  - `multiplier` (e.g., 3.0).  
- **Use Cases & Patterns**:  
  - **Trend detection**: SuperTrend line flips above/below price to signal bullish/bearish direction.  
  - Serves as a **trailing stop** mechanism that adapts to volatility.

---

## 19. T3 (Tillson’s T3 Moving Average)

- **Origin**: Developed by Tim Tillson.  
- **Description**: An advanced smoothing method that layers multiple EMAs internally and applies a “volume factor” to reduce lag while preserving smoothness.  
- **Common Parameters**:  
  - `period` (e.g., 14), `volumeFactor` (often 0.7).  
- **Use Cases & Patterns**:  
  - **Less lag** than a standard EMA, and can reduce whipsaws in sideways markets.  
  - Tunable “volume factor” lets traders set how aggressively T3 reacts to price changes.