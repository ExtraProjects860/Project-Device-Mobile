import smtplib
import logging
import os
from dotenv import load_dotenv

load_dotenv()

def InitializeMail ():
    if not (e_mail := os.getenv('EMAIL_USER')):
        raise ValueError("E-mail is not found")
    if not (password := os.getenv('EMAIL_PASSWORD')):
        raise ValueError("Password is not found")
    try:
        mailServer = smtplib.SMTP('smtp.gmail.com', 587)
        mailServer.starttls() 
        if not mailServer:
            raise ValueError("SMTP server is not found")
        mailServer.login(e_mail, password)
        return mailServer
    except Exception as e:
        logging.critical(f"Failed to connection to e-mail, error: {e}")