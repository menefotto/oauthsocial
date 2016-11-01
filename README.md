# OAUTHSOCIAL

Oauth social is a little package with some of the boiler plate associated with
implementing social login into an auth system. It provides and object and 
SocialStruct which takes an oauth2.Confing and query url and a function with the
signature func(response []byte) error, this func is called when the auth process
went well and the response is a json []byte containing the infos about the user.
It's tested and works. Tough I might not suits your need, you may consider it and
use it as a seed project for some other auth system.

NO TESTED ON A PRODUCTION SYSTEM.

