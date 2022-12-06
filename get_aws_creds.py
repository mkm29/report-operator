#!/usr/bin/env python

def get_profile_credentials(profile_name = 'trivy'):
    from configparser import ConfigParser
    from configparser import ParsingError
    from configparser import NoOptionError
    from configparser import NoSectionError
    from os import path
    config = ConfigParser()
    config.read([path.join(path.expanduser("~"),'.aws/credentials')])
    try:
        aws_access_key_id = config.get(profile_name, 'aws_access_key_id')
        aws_secret_access_key = config.get(profile_name, 'aws_secret_access_key')
    except ParsingError:
        print('Error parsing config file')
        raise
    except (NoSectionError, NoOptionError):
        print('Error reading config file')
        raise
    return aws_access_key_id, aws_secret_access_key

if __name__ == '__main__':
    aws_access_key_id, aws_secret_access_key = get_profile_credentials()
    print(f"{aws_access_key_id},{aws_secret_access_key}")