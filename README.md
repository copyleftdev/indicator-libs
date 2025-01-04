# Indicator Libs

A **Go (Golang) library** of popular technical analysis indicators. This is a **private** repository, providing a range of momentum and trend-based tools that can be integrated into quantitative trading or analysis workflows.

## Table of Contents

1. [Indicators Overview](#indicators-overview)  
   - [SMA (Simple Moving Average)](#1-sma-simple-moving-average)  
   - [EMA (Exponential Moving Average)](#2-ema-exponential-moving-average)  
   - [RSI (Relative Strength Index)](#3-rsi-relative-strength-index)  
   - [MACD (Moving Average Convergence/Divergence)](#4-macd-moving-average-convergence-divergence)  
   - [Bollinger Bands](#5-bollinger-bands)  
   - [Stochastic Oscillator](#6-stochastic-oscillator)  
   - [ATR (Average True Range)](#7-atr-average-true-range)  
   - [ADX (Average Directional Index)](#8-adx-average-directional-index)  
   - [CCI (Commodity Channel Index)](#9-cci-commodity-channel-index)  
   - [Williams %R](#10-williams-r)  
   - [OBV (On-Balance Volume)](#11-obv-on-balance-volume)  
   - [MFI (Money Flow Index)](#12-mfi-money-flow-index)  
   - [Ultimate Oscillator (UO)](#13-ultimate-oscillator-uo)  

---

## Indicators Overview

### 1. SMA (Simple Moving Average)
- **Origin**: A foundational smoothing technique in time-series analysis, used for decades in finance.  
- **Description**: Takes the arithmetic mean of prices over a fixed window.  
- **Common Parameters**:  
  - `window` (e.g., 20).  
- **Use Cases & Patterns**:  
  - **Trend identification** by smoothing data over time.  
  - **Crossover strategies** when used with multiple SMAs.  

### 2. EMA (Exponential Moving Average)
- **Origin**: A refinement of SMA, giving more weight to recent data points.  
- **Description**: Reduces lag by exponentially weighting recent prices.  
- **Common Parameters**:  
  - `window` (e.g., 20).  
- **Use Cases & Patterns**:  
  - **Faster reaction** to sudden price changes.  
  - Used in **MACD** calculations and multi-EMA cross strategies.

### 3. RSI (Relative Strength Index)
- **Origin**: Developed by J. Welles Wilder Jr. in the late 1970s.  
- **Description**: Ranges from 0 to 100, identifying overbought (>70) or oversold (<30) market conditions.  
- **Common Parameters**:  
  - `window` (e.g., 14).  
- **Use Cases & Patterns**:  
  - Spot possible **reversals**; look for RSI crossing key thresholds.  
  - **Divergence** between RSI and price can indicate momentum shifts.

### 4. MACD (Moving Average Convergence/Divergence)
- **Origin**: Created by Gerald Appel (1970s).  
- **Description**: Uses two EMAs (fast & slow) and a signal line to identify momentum and potential crossovers.  
- **Common Parameters**:  
  - `fastPeriod` (e.g., 12), `slowPeriod` (e.g., 26), `signalPeriod` (9).  
- **Use Cases & Patterns**:  
  - **Crossover** signals between MACD line & signal line.  
  - **Histogram** expansions to judge trend strength.

### 5. Bollinger Bands
- **Origin**: Developed by John Bollinger in the early 1980s.  
- **Description**: Plots an SMA-based middle band plus upper/lower bands at a specified number of standard deviations.  
- **Common Parameters**:  
  - `window` (e.g., 20), `num_std` (commonly 2).  
- **Use Cases & Patterns**:  
  - **Volatility** assessment (width of bands).  
  - **Bollinger Squeeze** signals potential breakouts.

### 6. Stochastic Oscillator
- **Origin**: Created by George C. Lane in the 1950s.  
- **Description**: Compares current close to the recent range, producing %K and %D lines typically in [0,100].  
- **Common Parameters**:  
  - `k_period` (14), `d_period` (3).  
- **Use Cases & Patterns**:  
  - Identifying **overbought/oversold** conditions.  
  - **%K–%D crossovers** for entry/exit signals.

### 7. ATR (Average True Range)
- **Origin**: By J. Welles Wilder Jr., also in the late 1970s.  
- **Description**: Measures volatility by considering the full price range and gaps.  
- **Common Parameters**:  
  - `window` (e.g., 14).  
- **Use Cases & Patterns**:  
  - Setting stop losses or position sizing based on **market volatility**.  
  - Filtering out low-volatility periods.

### 8. ADX (Average Directional Index)
- **Origin**: Also by Wilder in the late 1970s.  
- **Description**: Measures trend strength on a 0–100 scale, often paired with +DI and -DI.  
- **Common Parameters**:  
  - `window` (e.g., 14).  
- **Use Cases & Patterns**:  
  - Distinguishing **strong trending** markets (ADX > 25).  
  - +DI/-DI **crossovers** for bullish/bearish signals.

### 9. CCI (Commodity Channel Index)
- **Origin**: Invented by Donald Lambert (1980).  
- **Description**: Shows how far the current price is from its “average” over time; can move above +100 or below -100.  
- **Common Parameters**:  
  - `window` (often 14 or 20).  
- **Use Cases & Patterns**:  
  - Overbought/oversold detection outside ±100 range.  
  - **Divergence** signals potential trend shifts.

### 10. Williams %R
- **Origin**: Created by Larry Williams.  
- **Description**: Similar to Stochastic, ranges from 0 to -100, indicating close’s position relative to recent highs/lows.  
- **Common Parameters**:  
  - `window` (e.g., 14).  
- **Use Cases & Patterns**:  
  - Short-term **overbought/oversold** triggers.  
  - Often used to confirm momentum changes.

### 11. OBV (On-Balance Volume)
- **Origin**: Joseph Granville (1960s).  
- **Description**: Cumulative running total that adds volume on up days, subtracts on down days.  
- **Common Parameters**:  
  - Uses daily volume and close price movement.  
- **Use Cases & Patterns**:  
  - **Volume-based divergences** (OBV diverges from price).  
  - Identifying underlying strength/weakness of a trend.

### 12. MFI (Money Flow Index)
- **Origin**: Created by Gene Quong and Avrum Soudack; sometimes called a “volume-weighted RSI.”  
- **Description**: Ranges 0–100, factoring in both price and volume to detect overbought/oversold conditions.  
- **Common Parameters**:  
  - `window` (often 14).  
- **Use Cases & Patterns**:  
  - More sensitive than RSI (due to volume).  
  - Divergences often signal turning points.

### 13. Ultimate Oscillator (UO)
- **Origin**: By Larry Williams in 1976.  
- **Description**: Combines short-, medium-, and long-term price action into one oscillator (0–100).  
- **Common Parameters**:  
  - Typically uses periods 7, 14, 28 with weighting (4,2,1).  
- **Use Cases & Patterns**:  
  - Attempts to reduce false divergences by analyzing multiple timeframes.  
  - Overbought near 70–80, oversold near 20–30.

---

**Note**: Each indicator file in `indicators/` implements a constructor (e.g., `NewRSI(...)`) and a `Calculate(...)` method that returns the indicator’s values.
