from abc import ABCMeta, abstractmethod

class Payment(metaclass=ABCMeta):
    @abstractmethod
    def pay(self, money):
        pass


class AliPay(Payment):
    def pay(self, money):
        print(f'支付宝支付{money}元.')


class WechatPay(Payment):
    def pay(self, money):
        print(f'微信支付{money}元.')


class PaymentFactory:
    def create_factory(self, method):
        if method == 'alipay':
            return AliPay()
        else:
            return WechatPay()

# client
pf = PaymentFactory()
p = pf.create_factory('alipay')
p.pay(100)
