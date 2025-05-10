

# 🎮 GOitch: Your Twitch Adventure in the Terminal! 🚀

<p align="center">
  A snappy CLI Twitch client that brings streams and chats to your fingertips! <br/>
  <i>Low latency, high vibes, no browser required!</i>
</p>

---

# ⚙️ prerequisites 
- a media player VLC or MPV
- [Streamlink](https://github.com/streamlink/streamlink)

---
 
# 📦 Installation: Get Streaming in a Snap!
Ready to dive in? Install GOitch with one simple command:

```sh
go instal github.com/oalabsi4/goitch
```
---

## 🛠️ Setup: Connect GOitch to Twitch

To use GOitch, you’ll need to register an application with Twitch to get your `Client ID` and `Client Secret`. Follow these steps:

1. **Register Your App**:
   - Visit the [Twitch Developer Console](https://dev.twitch.tv/console/apps).
   - Click **Register Your Application**.
2. **Configure Your App**:
   - Set the **OAuth Redirect URL** to: `http://localhost:8080/oauth/callback`.
   - Choose **Category**: `Application Integration`.
   - Set **Client Type**: `Confidential`.
3. **Get Your Credentials**:
   - After registering, copy your **Client ID** and **Client Secret**.
4. **Create a `.env` File**:
   - In your GOitch project directory, create a file named `.env`.
   - Add your credentials like this:
     ```txt
     TWCLIENT=your_client_id
     TWSECRET=your_client_secret
     ```
   - Replace `your_client_id` and `your_client_secret` with the values from Twitch.

---

## 🎉 Features: Why GOitch Rocks!

- **Low-Latency Streaming**: Watch streams with minimal delay—feel the action in real-time! ⚡
- **Chat Like a Pro**: Join the Twitch chat, and vibe with the community. 💬
- **Light as a Feather**: No bloated browser, just a sleek CLI that’s fast and efficient. 🕊️


---

# 🔗 Sources

- [charmbraclet suit of tui tools](https://github.com/charmbracelet/bubbles)
- [Streamlink](https://github.com/streamlink/streamlink)
- [helix api](github.com/nicklaw5/helix )
